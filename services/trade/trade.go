package trade

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/quanghuy219/bots/config"
	"github.com/quanghuy219/bots/libs/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etherCommon "github.com/ethereum/go-ethereum/common"
)

func BuildTx() (*types.Transaction, error) {
	swapRouter, err := contracts.NewSwapRouterTransactor(etherCommon.HexToAddress(config.Cfg.SwapRouter), nil)
	if err != nil {
		return nil, err
	}

	amountIn := convert.MustFloatToWei(config.Cfg.AmountIn, 18)
	params := contracts.ISwapRouterExactInputSingleParams{
		TokenIn:   etherCommon.HexToAddress(config.Cfg.TokenIn),
		TokenOut:  etherCommon.HexToAddress(config.Cfg.TokenOut),
		Fee:       big.NewInt(int64(config.Cfg.PoolFee)),
		Recipient: etherCommon.HexToAddress(config.Cfg.Recipient),
		AmountIn:  amountIn,
	}

	return swapRouter.ExactInputSingle(&bind.TransactOpts{NoSend: true}, params)
}
