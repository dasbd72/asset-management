package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/models"
	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/rest"
	"github.com/dasbd72/go-exchange-sdk/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	c := rest.NewClient(cfg.BitfinexApiKey, cfg.BitfinexApiSecret)

	{
		res, err := c.GetTradingPairBooks(ctx, models.NewGetBooksRequest("tBTCUSD", models.BookPrecisionP0).Len(1))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetTradingPairBooks\n%s\n\n", string(b))
		// Test methods
		res[0].GetAmount()
		res[0].GetCount()
		res[0].GetPrice()
	}
	{
		res, err := c.GetFundingCurrencyBooks(ctx, models.NewGetBooksRequest("fUSD", models.BookPrecisionP0).Len(1))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetTradingPairBooks\n%s\n\n", string(b))
		// Test methods
		res[0].GetAmount()
		res[0].GetCount()
		res[0].GetPeriod()
		res[0].GetRate()
	}
}
