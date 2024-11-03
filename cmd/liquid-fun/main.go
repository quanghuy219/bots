package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/joho/godotenv"
	"github.com/quanghuy219/bots/common"
	"github.com/quanghuy219/bots/config"
	"github.com/quanghuy219/bots/services/gasprice"

	"github.com/ethereum/go-ethereum"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	API_URL                       = "https://api-test.liquid.fun/v1/swap/buildSwapTx"
	timeCoolDownEstimateGas int64 = 0
	bufferGasLimit                = 1.1
)

func main() {
	envFile := "config/.env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Error loading .env file from path %s, err %v", envFile, err)
	}

	config.InitConfig()

	ethClient, err := ethclient.Dial(config.Cfg.NodeEndpoint)
	if err != nil {
		log.Fatal("Fail to create ethclient: ", err)
	}

	publicKey, prvKey, err := common.GetAccountFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Public key:", publicKey.Hex())

	amountIn := convert.MustFloatToWei(config.Cfg.AmountIn, config.Cfg.TokenInDecimals)
	minDestAmount := convert.MustFloatToWei(config.Cfg.MinDestAmount, config.Cfg.TokenOutDecimals)

	metamaskGasPricer, err := gasprice.NewMetamaskGasPricer(config.Cfg.GasPriceEndpoint, nil)
	if err != nil {
		log.Fatal("Fail to create metamask gas pricer:", err)
	}
	cacheGasPricer := gasprice.NewCacheGasPricer(metamaskGasPricer, time.Second)

	tx, err := buildSwapTx(publicKey, amountIn, minDestAmount, cacheGasPricer)
	if err != nil {
		log.Fatal(err)
	}

	err = executeTx(ethClient, tx, publicKey, prvKey)
	if err != nil {
		log.Fatal(err)
	}
}

type BuildTxInput struct {
	ChainId     int
	Src         etherCommon.Address
	Dest        etherCommon.Address
	SrcAmount   *big.Int
	DestAmount  *big.Int
	UserAddress etherCommon.Address
}

type BuildTxOutput struct {
	Rates []*BuildSwapResult `json:"rates"`
}

type BuildSwapResult struct {
	TxObject *TxObject `json:"txObject"`
}

type TxObject struct {
	From  etherCommon.Address `json:"from"`
	To    etherCommon.Address `json:"to"`
	Value string              `json:"value"`
	Data  hexutil.Bytes       `json:"data"`
}

func getBuildSwapTx(userAddress etherCommon.Address, amountIn *big.Int, minDestAmount *big.Int) (*TxObject, error) {
	headers := map[string]string{
		"accept": "application/json",
	}

	queryParams := map[string]string{
		"chainId":       fmt.Sprintf("%d", config.Cfg.ChainId),
		"src":           config.Cfg.TokenIn,
		"dest":          config.Cfg.TokenOut,
		"srcAmount":     amountIn.String(),
		"minDestAmount": minDestAmount.String(),
		"userAddress":   userAddress.String(),
	}
	var resp BuildTxOutput
	err := common.MakeGetRequest(API_URL, headers, queryParams, 30*time.Second, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Rates) == 0 {
		return nil, fmt.Errorf("empty rates")
	}

	return resp.Rates[0].TxObject, nil
}

func buildSwapTx(userAddress etherCommon.Address, amountIn *big.Int, minDestAmount *big.Int, gasPricer gasprice.GasPricer) (*types.DynamicFeeTx, error) {
	txObject, err := getBuildSwapTx(userAddress, amountIn, minDestAmount)
	if err != nil {
		return nil, err
	}

	maxGasPriceGwei, gasTipCapGwei, err := gasPricer.GasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	maxGasPrice := convert.MustFloatToWei(maxGasPriceGwei, common.GweiDecimals)
	gasTipCap := convert.MustFloatToWei(config.Cfg.GasTipMultiplier*gasTipCapGwei, common.GweiDecimals)

	value := big.NewInt(0)
	if txObject.Value != "" {
		value, _ = new(big.Int).SetString(txObject.Value, 10)
	}

	tx := types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(config.Cfg.ChainId)),
		GasTipCap: gasTipCap,
		GasFeeCap: maxGasPrice,
		To:        &txObject.To,
		Data:      txObject.Data,
		Value:     value,
	}

	return &tx, nil
}

func executeTx(ethClient *ethclient.Client, tx *types.DynamicFeeTx, address etherCommon.Address, prvKey *ecdsa.PrivateKey) error {
	chainID := big.NewInt(int64(config.Cfg.ChainId))

	nonce, err := ethClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	tx.Nonce = nonce

	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(types.NewTx(tx), signer, prvKey)
	if err != nil {
		return err
	}

	var gasLimit uint64
	for {
		gasLimit, err = ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  address,
			To:    tx.To,
			Value: tx.Value,
			Data:  tx.Data,
		})
		if err != nil {
			log.Printf("error estimate gas: %v", err)
			if timeCoolDownEstimateGas > 0 {
				time.Sleep(time.Duration(timeCoolDownEstimateGas) * time.Second)
			}
			continue
		}

		break
	}

	bufferedGasLimit := uint64(float64(gasLimit) * bufferGasLimit)
	tx.Gas = bufferedGasLimit

	// Send transaction
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	timeCoolDownWaitForTx := time.Second
	timeWaitForTx := 10 * time.Second

	// Wait for transaction to be mined
	receipt, err := common.WaitForTransactionReceipt(context.Background(), ethClient, signedTx.Hash(), timeCoolDownWaitForTx, timeWaitForTx)
	if err != nil {
		log.Printf("Fail to get transaction receipt: transactionHash=%v error=%v", signedTx.Hash(), err)
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Printf("Transaction failed: transactionHash=%v status=%v", signedTx.Hash(), receipt.Status)
		return errors.New("transaction failed")
	}

	log.Printf("Successfully submit transaction: inputAmount=%v transactionHash=%v", config.Cfg.AmountIn, signedTx.Hash())
	return nil
}
