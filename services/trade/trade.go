package trade

import (
	"context"
	"log"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	swapRouter, err := contracts.NewSwapRouter02Transactor(etherCommon.HexToAddress(config.Cfg.SwapRouter), ethClient)
	if err != nil {
		return nil, err
	}

	recipient := fromAddress
	if config.Cfg.Recipient != "" {
		recipient = etherCommon.HexToAddress(config.Cfg.Recipient)
	}

	params := contracts.IV3SwapRouterExactInputSingleParams{
		TokenIn:           etherCommon.HexToAddress(config.Cfg.TokenIn),
		TokenOut:          etherCommon.HexToAddress(config.Cfg.TokenOut),
		Fee:               big.NewInt(int64(config.Cfg.PoolFee)),
		Recipient:         recipient,
		AmountIn:          amountIn,
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

	maxGasPrice, gasTipCap, err := getGasPriceWei(gasPricer)
	if err != nil {
		log.Printf("Fail to get gas price: error=%v", err)
	} else {
		opts.GasTipCap = gasTipCap
		opts.GasFeeCap = maxGasPrice
	}

	return swapRouter.ExactInputSingle(opts, params)
}

func BuildUniversalSwap(ethClient *ethclient.Client, gasPricer gasprice.GasPricer, fromAddress etherCommon.Address, amountIn, minDestAmount *big.Int) (*types.Transaction, error) {

	swapTimes := 2

	swapRouterAddress := etherCommon.HexToAddress(config.Cfg.SwapRouter)
	swapRouter, err := contracts.NewUniversalRouterTransactor(swapRouterAddress, ethClient)
	if err != nil {
		return nil, err
	}

	recipient := fromAddress
	if config.Cfg.Recipient != "" {
		recipient = etherCommon.HexToAddress(config.Cfg.Recipient)
	}

	opts := &bind.TransactOpts{
		NoSend: true,
		From:   fromAddress,
		Signer: func(a etherCommon.Address, t *types.Transaction) (*types.Transaction, error) {
			return t, nil
		},
		GasLimit: 1, // to skip estimate gas step
	}

	maxGasPrice, gasTipCap, err := getGasPriceWei(gasPricer)
	if err != nil {
		log.Printf("Fail to get gas price: error=%v", err)
	} else {
		opts.GasTipCap = gasTipCap
		opts.GasFeeCap = maxGasPrice
	}

	if err != nil {
		return nil, err
	}

	uint256Ty, _ := abi.NewType("uint256", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)
	addressTy, _ := abi.NewType("address", "", nil)
	addressArrTy, _ := abi.NewType("address[]", "", nil)

	v2SwapArgs := abi.Arguments{
		{Type: addressTy},
		{Type: uint256Ty},
		{Type: uint256Ty},
		{Type: addressArrTy},
		{Type: boolTy},
	}

	wrapArgs := abi.Arguments{
		{Type: addressTy},
		{Type: uint256Ty},
	}

	var inputs [][]byte
	var commands []byte
	if config.Cfg.IsNative {
		amountInNative := big.NewInt(0)
		amountInNative.Mul(amountIn, big.NewInt(int64(swapTimes)))

		opts.Value = amountInNative

		wrapByte, _ := wrapArgs.Pack(swapRouterAddress, amountInNative)
		commands = append(commands, 0x0b)
		inputs = append(inputs, wrapByte)
	}

	bytes, _ := v2SwapArgs.Pack(
		recipient,
		amountIn,
		minDestAmount,
		[]etherCommon.Address{etherCommon.HexToAddress(config.Cfg.TokenIn), etherCommon.HexToAddress(config.Cfg.TokenOut)},
		false,
	)
	for i := 0; i < swapTimes; i++ {
		commands = append(commands, 0x08)
		inputs = append(inputs, bytes)
	}
	return swapRouter.Execute(opts, commands, inputs)
}

func getGasPriceWei(gasPricer gasprice.GasPricer) (maxGasPrice *big.Int, gasTipCap *big.Int, err error) {
	maxGasPriceGwei, gasTipCapGwei, err := gasPricer.GasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}

	if config.Cfg.GasTipMultiplier > 1 {
		maxGasPriceGwei = maxGasPriceGwei * config.Cfg.GasTipMultiplier
		gasTipCapGwei = gasTipCapGwei * config.Cfg.GasTipMultiplier
	}

	maxGasPrice = convert.MustFloatToWei(maxGasPriceGwei, common.GweiDecimals)
	gasTipCap = convert.MustFloatToWei(gasTipCapGwei, common.GweiDecimals)

	return maxGasPrice, gasTipCap, nil
}
