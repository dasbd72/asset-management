package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	bitfinexRest "github.com/dasbd72/go-exchange-sdk/bitfinex/rest"
	"github.com/spf13/cobra"
)

func Bitfinex(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	// Create a new bitfinex client
	c := bitfinexRest.NewClient(
		os.Getenv("BFX_API_KEY"),
		os.Getenv("BFX_API_SECRET"),
	)

	// data, err := c.CallAPI(ctx, bfxRest.Request_builder{
	// 	Method:   http.MethodGet,
	// 	Endpoint: "/ticker",
	// 	SecType:  bfxRest.SecTypePrivate,
	// 	Params:   map[string]interface{}{},
	// }.Build())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(data))
	{
		data, err := c.GetWallets(ctx)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(b))
	}
	// {
	// 	data, err := c.GetFundingStats(ctx, "fUST")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for _, v := range data.FundingStatArray {
	// 		if v.MTS.Int() != data.FundingStatArray[0].MTS.Int() {
	// 			break
	// 		}
	// 		b, err := json.MarshalIndent(v, "", "  ")
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		log.Println(string(b))
	// 	}
	// }
	// {
	// 	data, err := c.GetActiveFundingOffers(ctx, "fUSD")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	b, err := json.MarshalIndent(data, "", "  ")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(string(b))
	// }
}
