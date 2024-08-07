package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	ChainId          int     `mapstructure:"chainId"`
	NodeEndpoint     string  `mapstructure:"nodeEndpoint"`
	SwapRouter       string  `mapstructure:"swapRouter"`
	StartTime        int64   `mapstructure:"startTime"`
	GasPriceEndpoint string  `mapstructure:"gasPriceEndpoint"`
	GasTipMultiplier float64 `mapstructure:"gasTipMultiplier"`
	Recipient        string  `mapstructure:"recipient"`
	TokenIn          string  `mapstructure:"tokenIn"`
	TokenOut         string  `mapstructure:"tokenOut"`
	PoolFee          int     `mapstructure:"poolFee"`
	AmountIn         float64 `mapstructure:"amountIn"`
	IsNative         bool    `mapstructure:"isNative"`
	MinDestAmount    float64 `mapstructure:"minDestAmount"`
}

func InitConfig() {
	cfg := &Config{}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("missing CONFIG_PATH env")
	}
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read viper config %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config %v", err)
	}
	Cfg = cfg
}
