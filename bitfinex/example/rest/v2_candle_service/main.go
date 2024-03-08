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
		res, err := c.GetCandlesRequest(ctx, models.NewGetCandlesRequest("trade:1m:tBTCUSD", models.CandleSTHist).Limit(1))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetCandlesRequest\n%s\n\n", string(b))
		// Test methods
		res[0].GetMTS()
		res[0].GetOpen()
		res[0].GetHigh()
		res[0].GetLow()
		res[0].GetClose()
		res[0].GetVolume()
	}
	{
		res, err := c.GetCandlesRequest(ctx, models.NewGetTradingPairsCandlesRequest(models.CandleTFFiveMinutes, "tBTCUSD", models.CandleSTHist).Limit(1))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetCandlesRequest\n%s\n\n", string(b))
		// Test methods
		res[0].GetMTS()
		res[0].GetOpen()
		res[0].GetHigh()
		res[0].GetLow()
		res[0].GetClose()
		res[0].GetVolume()
	}
	{
		res, err := c.GetCandlesRequest(ctx, models.NewGetFundingCurrenciesCandlesRequest(models.CandleTFOneMinute, "fUSD", 2, models.CandleSTHist).Limit(1))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetCandlesRequest\n%s\n\n", string(b))
		// Test methods
		res[0].GetMTS()
		res[0].GetOpen()
		res[0].GetHigh()
		res[0].GetLow()
		res[0].GetClose()
		res[0].GetVolume()
	}
	{
		res, err := c.GetCandlesRequest(ctx, models.NewGetCandlesRequest("trade:1m:fUSD:a10:p2:p10", models.CandleSTHist).Limit(1))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetCandlesRequest\n%s\n\n", string(b))
		// Test methods
		res[0].GetMTS()
		res[0].GetOpen()
		res[0].GetHigh()
		res[0].GetLow()
		res[0].GetClose()
		res[0].GetVolume()
	}
}
