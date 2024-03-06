package main

import (
	"context"
	"encoding/json"

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
		res, err := c.GetWalletStatus(ctx)
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		println(string(b))
	}
	{
		res, err := c.GetUserWalletBalance(ctx)
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res[0], "", "  ")
		if err != nil {
			panic(err)
		}
		println(string(b))
	}
}
