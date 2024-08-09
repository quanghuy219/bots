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
	config.InitConfig()

	err := handle()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
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

	maxPercent := 0.15
	if config.Cfg.LeftoverMaxPercent > 0 {
		maxPercent = config.Cfg.LeftoverMaxPercent
	}
	maxAmountIn := convert.MustFloatToWei(config.Cfg.AmountIn*maxPercent, 18)
	minDestMaxAmount := convert.MustFloatToWei(config.Cfg.MinDestAmount*maxPercent, 18)

	// 0.1%
	minPercent := 0.001
	minAmountIn := convert.MustFloatToWei(config.Cfg.AmountIn*minPercent, 18)
	minDestMinAmount := convert.MustFloatToWei(config.Cfg.MinDestAmount*minPercent, 18)

	var pivotAmountIn, pivotMinDestAmount *big.Int

	var successAmountIn, successMinDestAmount *big.Int

	maxTry := 10
	if config.Cfg.MaxTry > 0 {
		maxTry = config.Cfg.MaxTry
	}

	var successTx *types.Transaction
	// Binary search, only work if enough balance + approval for swap amount + gas fee
	for i := 0; i < maxTry; i++ {
		pivotAmountIn = new(big.Int).Add(maxAmountIn, minAmountIn)
		pivotAmountIn = new(big.Int).Quo(pivotAmountIn, big.NewInt(2))

		pivotMinDestAmount = new(big.Int).Add(minDestMaxAmount, minDestMinAmount)
		pivotMinDestAmount = new(big.Int).Quo(pivotMinDestAmount, big.NewInt(2))

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

		minAmountIn = pivotAmountIn
		minDestMinAmount = pivotMinDestAmount
	}

	if successTx == nil {
		return fmt.Errorf("=== Can not find leftover amount")
	} else {
		fmt.Printf("== Found swap option for amount in %.5f, minDestAmount: %.5f\n",
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
