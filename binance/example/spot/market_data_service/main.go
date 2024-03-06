package main

import (
	"context"
	"encoding/json"

	"github.com/dasbd72/go-exchange-sdk/binance/pkg/models"
	"github.com/dasbd72/go-exchange-sdk/binance/pkg/spot"
	"github.com/dasbd72/go-exchange-sdk/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	c := spot.NewClient(cfg.BinanceApiKey, cfg.BinanceApiSecret)

	{
		_, err := c.GetPing(ctx)
		if err != nil {
			panic(err)
		}
	}
	{
		res, err := c.GetServerTime(ctx)
		if err != nil {
			panic(err)
		}
		println(res.ServerTime)
	}
	{
		res, err := c.GetExchangeInfo(ctx)
		if err != nil {
			panic(err)
		}
		println(res.Symbols[0].Symbol)
	}
	{
		res, err := c.GetOrderBook(ctx, models.NewGetOrderBookRequest("BTCUSDT"))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res.Asks[0], "", "  ")
		if err != nil {
			panic(err)
		}
		println(string(b))
	}
	{
		res, err := c.GetAveragePrice(ctx, models.NewGetAveragePriceRequest("BTCUSDT"))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		println(string(b))
	}
}
