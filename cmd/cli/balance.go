package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dasbd72/asset-management/binance"
	"github.com/dasbd72/asset-management/max"
	"github.com/dasbd72/asset-management/okx"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func Balance(cmd *cobra.Command, args []string) error {
	// Load environment variables
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		return err
	}

	totalBalance := 0.0
	err = func() error {
		apiKey := os.Getenv("BINANCE_API_KEY")
		apiSecret := os.Getenv("BINANCE_API_SECRET")
		if apiKey == "" || apiSecret == "" {
			return nil
		}
		// Create a new Binance client
		c := binance.NewClient(apiKey, apiSecret)

		sum := 0.0
		wallet, err := c.GetUserWalletBalance(ctx)
		if err != nil {
			return err
		}
		for _, w := range *wallet {
			sum += w.Balance.Float64()
		}

		averagePrice, err := c.GetAveragePrice(ctx, &binance.GetAveragePriceRequest{
			Symbol: "BTCUSDT",
		})
		if err != nil {
			return err
		}
		btcPrice := averagePrice.Price.Float64()

		totalBalance += sum * btcPrice
		if useLog {
			log.Printf("Binance total balance: %.2f USDT\n", sum*btcPrice)
		}
		return nil
	}()
	if err != nil {
		return err
	}
	err = func() error {
		apiKey := os.Getenv("OKX_API_KEY")
		apiSecret := os.Getenv("OKX_API_SECRET")
		passphrase := os.Getenv("OKX_PASSPHRASE")
		if apiKey == "" || apiSecret == "" || passphrase == "" {
			return nil
		}
		// Create a new OKX client
		c := okx.NewClient(apiKey, apiSecret, passphrase)

		sum := 0.0
		wallet, err := c.GetBalance(ctx, &okx.GetBalanceRequest{})
		if err != nil {
			return err
		}
		for _, w := range wallet.Balances {
			sum += w.TotalEq.Float64()
		}
		funding, err := c.GetFundingBalances(ctx, &okx.GetFundingBalancesRequest{})
		if err != nil {
			return err
		}
		for _, f := range funding.Balances {
			sum += f.Bal.Float64()
		}
		savings, err := c.GetSavingBalance(ctx, &okx.GetSavingBalanceRequest{})
		if err != nil {
			return err
		}
		for _, s := range savings.Balances {
			price := 1.0
			if s.Ccy != "USDT" {
				ticker, err := c.GetTicker(ctx, &okx.GetTickerRequest{
					InstID: s.Ccy + "-USDT",
				})
				if err != nil {
					return err
				}
				price = ticker.Tickers[0].Last.Float64()
			}
			sum += s.Amt.Float64() * price
		}

		totalBalance += sum
		if useLog {
			log.Printf("OKX total balance: %.2f USDT\n", sum)
		}
		return nil
	}()
	if err != nil {
		return err
	}
	usdtToTWD, err := max.GetUsdtToTWD()
	if err != nil {
		return err
	}
	fmt.Printf("Total balance: %10.2f USDT\n", totalBalance)
	fmt.Printf("Total balance: %10.2f TWD\n", totalBalance*usdtToTWD)
	fmt.Printf("USDT to TWD: %.2f\n", usdtToTWD)

	return nil
}
