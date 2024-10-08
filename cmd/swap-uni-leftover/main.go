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
	timeCoolDownEstimateGas = 0
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

	nonce, err := ethClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	maxPercent := 0.15
	if config.Cfg.LeftoverMaxPercent > 0 {
		maxPercent = config.Cfg.LeftoverMaxPercent
	}
	maxAmountIn := convert.MustFloatToWei(config.Cfg.AmountIn*maxPercent, 18)
	minDestMaxAmount := convert.MustFloatToWei(config.Cfg.MinDestAmount*maxPercent, 18)

	// 0.1%
	minPercent := 0.001
	if config.Cfg.LeftoverMinPercent > 0 {
		minPercent = config.Cfg.LeftoverMinPercent
	}
	minAmountIn := convert.MustFloatToWei(config.Cfg.AmountIn*minPercent, 18)
	minDestMinAmount := convert.MustFloatToWei(config.Cfg.MinDestAmount*minPercent, 18)

	var pivotAmountIn, pivotMinDestAmount *big.Int

	var successAmountIn, successMinDestAmount *big.Int

	maxTry := 10
	if config.Cfg.MaxTry > 0 {
		maxTry = config.Cfg.MaxTry
	}

	diffThreshold := 0.0000001
	if config.Cfg.DiffThreshold > 0 {
		diffThreshold = config.Cfg.DiffThreshold
	}

	log.Printf("Start finding leftover from %.2f to %.2f, diff threshold is %.10f, max try %d\n",
		maxPercent, minPercent, diffThreshold, maxTry,
	)

	diffThresholdWei := convert.MustFloatToWei(diffThreshold, 18)

	var successTx *types.Transaction
	// Binary search, only work if enough balance + approval for swap amount + gas fee
	for i := 0; i < maxTry; i++ {
		// if diff below threshold => break
		diffAmountIn := new(big.Int).Sub(maxAmountIn, minAmountIn)
		if diffAmountIn.Cmp(diffThresholdWei) < 0 {
			log.Println("Diff is below threshold", diffThresholdWei, diffAmountIn)
			break
		}

		pivotAmountIn = new(big.Int).Add(maxAmountIn, minAmountIn)
		pivotAmountIn = new(big.Int).Quo(pivotAmountIn, big.NewInt(2))

		pivotMinDestAmount = new(big.Int).Add(minDestMaxAmount, minDestMinAmount)
		pivotMinDestAmount = new(big.Int).Quo(pivotMinDestAmount, big.NewInt(2))

		log.Println("Trying", i, pivotAmountIn, pivotMinDestAmount)
		tx, err := trade.BuildTx(ethClient, gasPricer, address, pivotAmountIn, pivotMinDestAmount)
		// can not swap
		if err != nil {
			log.Printf("error buildTx: %v", err)
			maxAmountIn = pivotAmountIn
			minDestMaxAmount = pivotMinDestAmount
			continue
		}

		_, err = ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  address,
			To:    tx.To(),
			Value: tx.Value(),
			Data:  tx.Data(),
		})

		// can not swap
		if err != nil {
			log.Printf("error estimate gas: %v", err)
			maxAmountIn = pivotAmountIn
			minDestMaxAmount = pivotMinDestAmount
			continue
		}

		successTx = tx
		successAmountIn = pivotAmountIn
		successMinDestAmount = pivotMinDestAmount

		log.Printf("Got amount In %.8f\n", convert.WeiToFloat(successAmountIn, 18))

		minAmountIn = pivotAmountIn
		minDestMinAmount = pivotMinDestAmount
	}

	if successTx == nil {
		return fmt.Errorf("can not find leftover amount")
	} else {
		log.Printf("Found swap option for amount in %.8f, minDestAmount: %.8f\n",
			convert.WeiToFloat(successAmountIn, 18), convert.WeiToFloat(successMinDestAmount, 18))
	}

	var gasLimit uint64
	for {
		gasLimit, err = ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
			From:  address,
			To:    successTx.To(),
			Value: successTx.Value(),
			Data:  successTx.Data(),
		})
		if err != nil {
			log.Printf("error estimate gas: %v", err)
			if timeCoolDownEstimateGas > 0 {
				time.Sleep(timeCoolDownEstimateGas * time.Second)
			}

			continue
		}

		break
	}

	bufferedGasLimit := uint64(float64(gasLimit) * bufferGasLimit)
	rawTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(config.Cfg.ChainId)),
		Nonce:     nonce,
		GasTipCap: successTx.GasTipCap(),
		GasFeeCap: successTx.GasFeeCap(),
		Gas:       bufferedGasLimit,
		To:        successTx.To(),
		Value:     successTx.Value(),
		Data:      successTx.Data(),
	}

	signedTx, err := types.SignNewTx(prvKey, types.LatestSignerForChainID(big.NewInt(int64(config.Cfg.ChainId))), rawTx)
	if err != nil {
		return err
	}

	log.Printf("Submit transaction: inputAmount=%v\ntransactionHash=%v", successAmountIn, signedTx.Hash())
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

	log.Printf("Successfully submit transaction: inputAmount=%v transactionHash=%v", successAmountIn, signedTx.Hash())

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
