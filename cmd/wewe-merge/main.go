package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/quanghuy219/bots/common"
	"github.com/quanghuy219/bots/config"
	"github.com/quanghuy219/bots/libs/contracts"
	"github.com/quanghuy219/bots/services/gasprice"
)

const (
	bufferGasLimit          = 1.2
	timeWaitForTx           = 10 * time.Second
	timeCoolDownWaitForTx   = time.Second
	timeCoolDownEstimateGas = time.Second
)

var (
	weweAddress      = etherCommon.HexToAddress("0x6b9bb36519538e0c073894e964e90172e1c0b41f")
	vultMergeAddress = etherCommon.HexToAddress("0x30091c97Fd47873c44D03F7F1b960473D300D269")
)

func main() {
	envFile := "config/.env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Error loading .env file from path %s, err %v", envFile, err)
	}

	config.InitConfig()

	err = handle()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done.")
}

func handle() error {
	delay := time.Until(time.Unix(config.Cfg.StartTime, 0))
	if delay > 0 {
		log.Printf("Wait %v before starting to make trades\n", delay)
		time.Sleep(delay)
	}

	ethClient, err := ethclient.Dial(config.Cfg.NodeEndpoint)
	if err != nil {
		log.Println("Fail to create ethclient: ", err)
		return err
	}

	metamaskGasPricer, err := gasprice.NewMetamaskGasPricer(config.Cfg.GasPriceEndpoint, nil)
	if err != nil {
		log.Println("Fail to create metamask gas pricer:", err)
		return err
	}
	cacheGasPricer := gasprice.NewCacheGasPricer(metamaskGasPricer, time.Second)

	return makeTrade(ethClient, cacheGasPricer)
}

func makeTrade(ethClient *ethclient.Client, gasPricer gasprice.GasPricer) error {
	address, prvKey, err := getKey()
	if err != nil {
		return err
	}

	weweContract, _ := contracts.NewWeweCaller(weweAddress, ethClient)
	amountIn, err := weweContract.BalanceOf(nil, address)
	if err != nil {
		log.Printf("err get balance %v", err)
		return err
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	tx, err := buildTx(ethClient, address, amountIn, gasPricer)
	if err != nil {
		return err
	}

	var gasLimit uint64
	for {
		gasLimit, err = ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  address,
			To:    tx.To(),
			Value: tx.Value(),
			Data:  tx.Data(),
		})
		if err != nil {
			log.Printf("error estimate gas: %v", err)
			time.Sleep(timeCoolDownEstimateGas)
			continue
		} else {
			break
		}
	}

	bufferedGasLimit := uint64(float64(gasLimit) * bufferGasLimit)
	rawTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(config.Cfg.ChainId)),
		Nonce:     nonce,
		GasTipCap: tx.GasTipCap(),
		GasFeeCap: tx.GasFeeCap(),
		Gas:       bufferedGasLimit,
		To:        tx.To(),
		Value:     tx.Value(),
		Data:      tx.Data(),
	}

	signedTx, err := types.SignNewTx(prvKey, types.LatestSignerForChainID(big.NewInt(int64(config.Cfg.ChainId))), rawTx)
	if err != nil {
		return err
	}

	log.Printf("Submit transaction: inputAmount=%v\ntransactionHash=%v", config.Cfg.AmountIn, signedTx.Hash())
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("error send transaction %v", err)
		return err
	}

	// Wait for transaction to be mined
	receipt, err := waitForTransactionReceipt(context.Background(), ethClient, signedTx.Hash(), timeWaitForTx)
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

func buildTx(
	ethClient *ethclient.Client,
	sender etherCommon.Address,
	amountIn *big.Int,
	gasPricer gasprice.GasPricer,
) (*types.Transaction, error) {
	weweCaller, _ := contracts.NewWeweTransactor(weweAddress, ethClient)
	opts := &bind.TransactOpts{
		From: sender,
		Signer: func(a etherCommon.Address, t *types.Transaction) (*types.Transaction, error) {
			return t, nil
		},
		GasLimit: 1,
		NoSend:   true,
	}
	maxGasPriceGwei, gasTipCapGwei, err := gasPricer.GasPrice(context.Background())
	if err != nil {
		log.Printf("Fail to get gas price: error=%v", err)
	} else {
		maxGasPrice := convert.MustFloatToWei(maxGasPriceGwei, common.GweiDecimals)
		gasTipCap := convert.MustFloatToWei(config.Cfg.GasTipMultiplier*gasTipCapGwei, common.GweiDecimals)
		opts.GasTipCap = gasTipCap
		opts.GasFeeCap = maxGasPrice
	}
	return weweCaller.ApproveAndCall(opts, vultMergeAddress, amountIn, []byte{})
}

func getKey() (etherCommon.Address, *ecdsa.PrivateKey, error) {
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return etherCommon.Address{}, nil, fmt.Errorf("failed to parse private key %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return etherCommon.Address{}, nil, errors.New("failed to get public Key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress, privateKey, nil
}

func waitForTransactionReceipt(ctx context.Context, ethClient *ethclient.Client, txHash etherCommon.Hash, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		receipt, err := ethClient.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, fmt.Errorf("error fetching receipt: %v", err)
		}

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("transaction not mined within %v", timeout)
		case <-ticker.C:
			time.Sleep(timeCoolDownWaitForTx)
			continue
		}
	}
}
