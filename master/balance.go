package master

import (
	"context"

	"github.com/dasbd72/asset-management/binance"
	"github.com/dasbd72/asset-management/max"
	"github.com/dasbd72/asset-management/okx"
)

type (
	Balance struct {
		Usdt float64
		Twd  float64
	}
)

func (c *Client) GetBalance(ctx context.Context) (*Balance, error) {
	var (
		totalBalanceUsdt float64
		totalBalanceTwd  float64
	)
	funcs := []func() error{
		func() error {
			sum := 0.0
			wallet, err := c.binanceClient.GetUserWalletBalance(ctx)
			if err != nil {
				return err
			}
			for _, w := range *wallet {
				sum += w.Balance.Float64()
			}

			averagePrice, err := c.binanceClient.GetAveragePrice(ctx, binance.NewGetAveragePriceRequest("BTCUSDT"))
			if err != nil {
				return err
			}
			btcPrice := averagePrice.Price.Float64()

			totalBalanceUsdt += sum * btcPrice
			return nil
		},
		func() error {
			sum := 0.0
			wallet, err := c.okxClient.GetBalance(ctx, okx.NewGetBalanceRequest())
			if err != nil {
				return err
			}
			for _, w := range wallet.Balances {
				sum += w.TotalEq.Float64()
			}
			funding, err := c.okxClient.GetFundingBalances(ctx, okx.NewGetFundingBalancesRequest())
			if err != nil {
				return err
			}
			for _, f := range funding.Balances {
				sum += f.Bal.Float64()
			}
			savings, err := c.okxClient.GetSavingBalance(ctx, okx.NewGetSavingBalanceRequest())
			if err != nil {
				return err
			}
			for _, s := range savings.Balances {
				price := 1.0
				if s.Ccy != "USDT" {
					ticker, err := c.okxClient.GetTicker(ctx, okx.NewGetTickerRequest(s.Ccy+"-USDT"))
					if err != nil {
						return err
					}
					price = ticker.Tickers[0].Last.Float64()
				}
				sum += s.Amt.Float64() * price
			}

			totalBalanceUsdt += sum
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
