package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/example/config"
	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/models"
	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/rest"
)

var (
	submit    bool
	cancel    bool
	id        int
	cancelAll bool
)

func init() {
	flag.BoolVar(&submit, "submit", false, "Use submit operation")
	flag.BoolVar(&cancel, "cancel", false, "Use cancel operation")
	flag.BoolVar(&cancelAll, "cancel-all", false, "Use cancel all operation")
	flag.IntVar(&id, "id", 0, "Offer ID")
	flag.Parse()
}

func main() {
	cfg := config.Load()

	ctx := context.Background()
	c := rest.NewClient(cfg.BitfinexApiKey, cfg.BitfinexApiSecret)

	if submit {
		res, err := c.SubmitFundingOffer(ctx, models.NewSubmitFundingOfferRequest(models.FRRLIMIT, "fUSD", "150", "0.00041", 2))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("SubmitFundingOffer\n%s\n\n", string(b))
	}
	if cancel {
		res, err := c.CancelFundingOffer(ctx, models.NewCancelFundingOfferRequest(id))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("CancelFundingOffer\n%s\n\n", string(b))
	}
	if cancelAll {
		res, err := c.CancelAllFundingOffer(ctx, models.NewCancelAllFundingOfferRequest().Currency("USD"))
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("CancelAllFundingOffers\n%s\n\n", string(b))
	}
	{
		res, err := c.GetFundingStats(ctx, "fUSD")
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetFundingStats\n%s\n\n", string(b))
	}
	{
		res, err := c.GetAllActiveFundingOffers(ctx)
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetAllActiveFundingOffers\n%s\n\n", string(b))
	}
	{
		res, err := c.GetActiveFundingOffers(ctx, "fUSD")
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetActiveFundingOffers\n%s\n\n", string(b))
	}
	{
		res, err := c.GetFundingOffersHistory(ctx, "fUSD")
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res.FundingOffers[0], "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetFundingOffersHistory\n%s\n\n", string(b))
	}
	{
		res, err := c.GetFundingInfo(ctx, "fUSD")
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetFundingInfo\n%s\n\n", string(b))
	}
}
