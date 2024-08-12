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
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/quanghuy219/bots/config"
	"github.com/quanghuy219/bots/services/gasprice"
	"github.com/quanghuy219/bots/services/trade"
)

const (
	bufferGasLimit          = 1.1
	timeWaitForTx           = 10 * time.Second
	timeCoolDownWaitForTx   = time.Second
	timeCoolDownEstimateGas = time.Second
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

	amountIn := convert.MustFloatToWei(config.Cfg.AmountIn, 18)
	minDestAmount := convert.MustFloatToWei(config.Cfg.MinDestAmount, 18)
	tx, err := trade.BuildTx(ethClient, gasPricer, address, amountIn, minDestAmount)
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

	nonce, err := ethClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
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
