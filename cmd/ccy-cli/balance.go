package main

import (
	"context"
	"fmt"

	binanceSpot "github.com/dasbd72/go-exchange-sdk/binance/pkg/spot"
	bitfinexRest "github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/rest"
	"github.com/dasbd72/go-exchange-sdk/config"
	"github.com/dasbd72/go-exchange-sdk/manager"
	okxRest "github.com/dasbd72/go-exchange-sdk/okx/pkg/rest"
	"github.com/spf13/cobra"
)

func Balance(cmd *cobra.Command, args []string) error {
	// Load environment variables
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if cfg.BinanceApiKey == "" || cfg.BinanceApiSecret == "" {
		return nil
	}
	if cfg.OKXApiKey == "" || cfg.OKXApiSecret == "" || cfg.OKXPassphrase == "" {
		return nil
	}
	if cfg.BitfinexApiKey == "" || cfg.BitfinexApiSecret == "" {
		return nil
	}

	c := manager.Client_builder{
		BinanceClient:      binanceSpot.NewClient(cfg.BinanceApiKey, cfg.BinanceApiSecret),
		OkxClient:          okxRest.NewClient(cfg.OKXApiKey, cfg.OKXApiSecret, cfg.OKXPassphrase),
		BitfinexRestClient: bitfinexRest.NewClient(cfg.BitfinexApiKey, cfg.BitfinexApiSecret),
	}.Build()

	// Get balance
	balanceUsdt := 0.0
	balanceTwd := 0.0
	// Get Binance balance
	binanceBalance, err := c.GetBinanceBalance(ctx)
	if err != nil {
		return err
	}
	balanceUsdt += binanceBalance
	// Get OKX balance
	okxBalance, err := c.GetOkxBalance(ctx)
	if err != nil {
		return err
	}
	balanceUsdt += okxBalance
	// Get Bitfinex balance
	bitfinexBalance, err := c.GetBitfinexBalance(ctx)
	if err != nil {
		return err
	}
	balanceUsdt += bitfinexBalance
	// Get USDT to TWD
	usdtToTWD, err := c.GetUsdtToTWD()
	if err != nil {
		return err
	}
	balanceTwd = balanceUsdt * usdtToTWD
	fmt.Printf("Total balance: %10.2f USDT\n", balanceUsdt)
	fmt.Printf("Total balance: %10.2f TWD\n", balanceTwd)
	fmt.Printf("USDT to TWD: %.2f\n", usdtToTWD)

	return nil
}
