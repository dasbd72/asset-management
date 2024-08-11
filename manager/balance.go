package manager

import (
	"context"
	"fmt"
	"log/slog"

	binanceModels "github.com/dasbd72/go-exchange-sdk/binance/pkg/models"
	"github.com/dasbd72/go-exchange-sdk/max"
	okxModels "github.com/dasbd72/go-exchange-sdk/okx/pkg/models"
)

type (
	Balance struct {
		Usdt float64
		Twd  float64
	}
)

// GetOkxSymbolPrice returns the price of the symbol in USDT
func (c *Client) GetOkxSymbolPrice(ctx context.Context, symbol string) (float64, error) {
	ticker, err := c.okxClient.GetTicker(ctx, okxModels.NewGetTickerRequest(symbol))
	if err != nil {
		return 0, err
	}
	if len(ticker.Tickers) == 0 {
		slog.Warn(fmt.Sprintf("no ticker found for %s", symbol))
		return 0, nil
	}
	return ticker.Tickers[0].Last.Float64(), nil
}

func (c *Client) GetBalance(ctx context.Context) (*Balance, error) {
	var (
		totalBalanceUsdt float64
		totalBalanceTwd  float64
	)

	getBinanceBalance := func() (float64, error) {
		if c.binanceClient == nil {
			// Skip if binance client is not set
			return 0, nil
		}
		sum := 0.0
		// Get balance from wallet
		wallet, err := c.binanceClient.GetUserWalletBalance(ctx)
		if err != nil {
			return 0, err
		}
		for _, w := range wallet {
			sum += w.Balance.Float64()
		}

		averagePrice, err := c.binanceClient.GetAveragePrice(ctx, binanceModels.NewGetAveragePriceRequest("BTCUSDT"))
		if err != nil {
			return 0, err
		}
		btcPrice := averagePrice.Price.Float64()

		balance := sum * btcPrice
		slog.Info(fmt.Sprintf("Binance balance: %f", balance))
		return balance, nil
	}

	getOkxBalance := func() (float64, error) {
		if c.okxClient == nil {
			// Skip if okx client is not set
			return 0, nil
		}
		balance := 0.0
		// Get balance from wallet
		wallet, err := c.okxClient.GetBalance(ctx, okxModels.NewGetBalanceRequest())
		if err != nil {
			return 0, err
		}
		for _, w := range wallet.Balances {
			for _, detail := range w.Details {
				price := 1.0
				if detail.Ccy != "USDT" {
					price, err = c.GetOkxSymbolPrice(ctx, detail.Ccy+"-USDT")
					if err != nil {
						return 0, err
					}
				}
				balance += detail.Eq.Float64() * price
			}
		}
		// Get balance from funding
		funding, err := c.okxClient.GetFundingBalances(ctx, okxModels.NewGetFundingBalancesRequest())
		if err != nil {
			return 0, err
		}
		for _, f := range funding.Balances {
			price := 1.0
			if f.Ccy != "USDT" {
				price, err = c.GetOkxSymbolPrice(ctx, f.Ccy+"-USDT")
				if err != nil {
					return 0, err
				}
			}
			balance += f.Bal.Float64() * price
		}
		// Get balance from saving
		savings, err := c.okxClient.GetSavingBalance(ctx, okxModels.NewGetSavingBalanceRequest())
		if err != nil {
			return 0, err
		}
		for _, s := range savings.Balances {
			price := 1.0
			if s.Ccy != "USDT" {
				price, err = c.GetOkxSymbolPrice(ctx, s.Ccy+"-USDT")
				if err != nil {
					return 0, err
				}
			}
			balance += s.Amt.Float64() * price
		}
		slog.Info(fmt.Sprintf("OKX balance: %f", balance))
		return balance, nil
	}

	getBitfinexBalance := func() (float64, error) {
		if c.bitfinexRestClient == nil {
			// Skip if bitfinex client is not set
			return 0, nil
		}
		balance := 0.0
		// Get balance from wallet
		res, err := c.bitfinexRestClient.GetWallets(ctx)
		if err != nil {
			return 0, err
		}
		for _, w := range res.Wallets {
			if w.Currency.String() == "USD" || w.Currency.String() == "UST" {
				balance += w.Balance.Float64()
			}
		}
		slog.Info(fmt.Sprintf("Bitfinex balance: %f", balance))
		return balance, nil
	}

	funcs := []func() error{
		func() error {
			balance, err := getBinanceBalance()
			if err != nil {
				return err
			}
			totalBalanceUsdt += balance
			return nil
		},
		func() error {
			balance, err := getOkxBalance()
			if err != nil {
				return err
			}
			totalBalanceUsdt += balance
			return nil
		},
		func() error {
			balance, err := getBitfinexBalance()
			if err != nil {
				return err
			}
			totalBalanceUsdt += balance
			return nil
		},
		func() error {
			usdtToTWD, err := max.GetUsdtToTWD()
			if err != nil {
				return err
			}
			totalBalanceTwd = totalBalanceUsdt * usdtToTWD
			return nil
		},
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			return nil, err
		}
	}
	return &Balance{
		Usdt: totalBalanceUsdt,
		Twd:  totalBalanceTwd,
	}, nil
}
