package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dasbd72/asset-management/binance"
	"github.com/dasbd72/asset-management/master"
	"github.com/dasbd72/asset-management/max"
	"github.com/dasbd72/asset-management/okx"
	"github.com/spf13/cobra"
)

func Balance(cmd *cobra.Command, args []string) error {
	// Load environment variables
	ctx := context.Background()

	binanceApiKey := os.Getenv("BINANCE_API_KEY")
	binanceApiSecret := os.Getenv("BINANCE_API_SECRET")
	if binanceApiKey == "" || binanceApiSecret == "" {
		return nil
	}

	okxApiKey := os.Getenv("OKX_API_KEY")
	okxApiSecret := os.Getenv("OKX_API_SECRET")
	okxPassphrase := os.Getenv("OKX_PASSPHRASE")
	if okxApiKey == "" || okxApiSecret == "" || okxPassphrase == "" {
		return nil
	}

	c := master.NewClient(
		binance.NewClient(binanceApiKey, binanceApiSecret),
		okx.NewClient(okxApiKey, okxApiSecret, okxPassphrase),
	)

	balance, err := c.GetBalance(ctx)
	if err != nil {
		return err
	}
	usdtToTWD, err := max.GetUsdtToTWD()
	if err != nil {
		return err
	}
	fmt.Printf("Total balance: %10.2f USDT\n", balance.Usdt)
	fmt.Printf("Total balance: %10.2f TWD\n", balance.Twd)
	fmt.Printf("USDT to TWD: %.2f\n", usdtToTWD)

	return nil
}
