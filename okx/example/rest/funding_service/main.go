package main

import (
	"context"
	"encoding/json"

	"github.com/dasbd72/go-exchange-sdk/config"
	"github.com/dasbd72/go-exchange-sdk/okx/pkg/models"
	"github.com/dasbd72/go-exchange-sdk/okx/pkg/rest"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	c := rest.NewClient(cfg.OKXApiKey, cfg.OKXApiSecret, cfg.OKXPassphrase)

	{
		res, err := c.GetFundingBalances(ctx, models.NewGetFundingBalancesRequest())
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
