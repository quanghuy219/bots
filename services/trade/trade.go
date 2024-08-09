package trade

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quanghuy219/bots/common"
	"github.com/quanghuy219/bots/config"
	"github.com/quanghuy219/bots/libs/contracts"
	"github.com/quanghuy219/bots/services/gasprice"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etherCommon "github.com/ethereum/go-ethereum/common"
)

func BuildTx(ethClient *ethclient.Client, gasPricer gasprice.GasPricer, fromAddress etherCommon.Address, amountIn, minDestAmount *big.Int) (*types.Transaction, error) {
	swapRouter, err := contracts.NewSwapRouterTransactor(etherCommon.HexToAddress(config.Cfg.SwapRouter), ethClient)
	if err != nil {
		return nil, err
	}

	deadline := time.Now().Add(time.Hour)
	recipient := fromAddress
	if config.Cfg.Recipient != "" {
		recipient = etherCommon.HexToAddress(config.Cfg.Recipient)
	}
	params := contracts.ISwapRouterExactInputSingleParams{
		TokenIn:           etherCommon.HexToAddress(config.Cfg.TokenIn),
		TokenOut:          etherCommon.HexToAddress(config.Cfg.TokenOut),
		Fee:               big.NewInt(int64(config.Cfg.PoolFee)),
		Recipient:         recipient,
		AmountIn:          amountIn,
		Deadline:          big.NewInt(deadline.Unix()),
		AmountOutMinimum:  minDestAmount,
		SqrtPriceLimitX96: big.NewInt(0),
	}

	opts := &bind.TransactOpts{
		NoSend: true,
		From:   fromAddress,
		Signer: func(a etherCommon.Address, t *types.Transaction) (*types.Transaction, error) {
			return t, nil
		},
		GasLimit: 1, // to skip estimate gas step
	}
	if config.Cfg.IsNative {
		opts.Value = amountIn
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

	return swapRouter.ExactInputSingle(opts, params)
}
