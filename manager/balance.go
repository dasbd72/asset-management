package manager

import (
	"context"
	"fmt"
	"log/slog"
	"math"

	binanceModels "github.com/dasbd72/go-exchange-sdk/binance/pkg/models"
	"github.com/dasbd72/go-exchange-sdk/max"
	okxModels "github.com/dasbd72/go-exchange-sdk/okx/pkg/models"
)

// GetBinanceBalance returns the balance of Binance in USDT
func (c *Client) GetBinanceBalance(ctx context.Context) (float64, error) {
	if c.binanceClient == nil {
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

	// Get average price of BTC
	averagePrice, err := c.binanceClient.GetAveragePrice(ctx, binanceModels.NewGetAveragePriceRequest("BTCUSDT"))
	if err != nil {
		return 0, err
	}
	btcPrice := averagePrice.Price.Float64()

	balance := sum * btcPrice
	slog.Info(fmt.Sprintf("Binance balance: %f", balance))
	return balance, nil
}

// getOkxSymbolPrice returns the price of the symbol in USDT
func (c *Client) getOkxSymbolPrice(ctx context.Context, symbol string) (float64, error) {
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

// getOkxJumpstartBalance returns the balance of Jumpstart by the net flow from wallet to Jumpstart
func (c *Client) getOkxJumpstartBalance(ctx context.Context) (float64, error) {
	balance := 0.0
	// Get balance from Jumpstart, from wallet to Jumpstart, hence minus
	outBills, err := c.okxClient.GetFundingBills(ctx, okxModels.NewGetFundingBillsRequest().Type("78"))
	if err != nil {
		return 0, err
	}
	for _, b := range outBills.Bills {
		price := 1.0
		if b.Ccy != "USDT" {
			price, err = c.getOkxSymbolPrice(ctx, b.Ccy+"-USDT")
			if err != nil {
				return 0, err
			}
		}
		balance -= b.BalChg.Float64() * price
	}
	// Get balance from Jumpstart, from Jumpstart to wallet, hence minus
	inBills, err := c.okxClient.GetFundingBills(ctx, okxModels.NewGetFundingBillsRequest().Type("124"))
	if err != nil {
		return 0, err
	}
	for _, b := range inBills.Bills {
		price := 1.0
		if b.Ccy != "USDT" {
			price, err = c.getOkxSymbolPrice(ctx, b.Ccy+"-USDT")
			if err != nil {
				return 0, err
			}
		}
		balance -= b.BalChg.Float64() * price
	}
	balance = math.Max(balance, 0)
	slog.Info(fmt.Sprintf("OKX Jumpstart balance: %f", balance))
	return balance, nil
}

// GetOkxBalance returns the balance of OKX in USDT
func (c *Client) GetOkxBalance(ctx context.Context) (float64, error) {
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
				price, err = c.getOkxSymbolPrice(ctx, detail.Ccy+"-USDT")
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
			price, err = c.getOkxSymbolPrice(ctx, f.Ccy+"-USDT")
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
			price, err = c.getOkxSymbolPrice(ctx, s.Ccy+"-USDT")
			if err != nil {
				return 0, err
			}
		}
		balance += s.Amt.Float64() * price
	}
	// Get balance from Jumpstart
	jumpstartBalance, err := c.getOkxJumpstartBalance(ctx)
	if err != nil {
		return 0, err
	}
	balance += jumpstartBalance
	slog.Info(fmt.Sprintf("OKX balance: %f", balance))
	return balance, nil
}

// GetBitfinexBalance returns the balance of Bitfinex in USDT
func (c *Client) GetBitfinexBalance(ctx context.Context) (float64, error) {
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

// GetUsdtToTWD returns the exchange rate of USDT to TWD
func (c *Client) GetUsdtToTWD() (float64, error) {
	return max.GetUsdtToTWD()
}

// GetBalance returns the balance of all exchanges in USDT
func (c *Client) GetBalance(ctx context.Context) (float64, error) {
	balanceUsdt := 0.0
	// Get Binance balance
	binanceBalance, err := c.GetBinanceBalance(ctx)
	if err != nil {
		return 0, err
	}
	balanceUsdt += binanceBalance
	// Get OKX balance
	okxBalance, err := c.GetOkxBalance(ctx)
	if err != nil {
		return 0, err
	}
	balanceUsdt += okxBalance
	// Get Bitfinex balance
	bitfinexBalance, err := c.GetBitfinexBalance(ctx)
	if err != nil {
		return 0, err
	}
	balanceUsdt += bitfinexBalance
	slog.Info(fmt.Sprintf("Total balance: %f USDT", balanceUsdt))
	return balanceUsdt, nil
}
