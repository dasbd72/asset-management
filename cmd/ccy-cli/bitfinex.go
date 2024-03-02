package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dasbd72/go-exchange-sdk/bitfinex"
	"github.com/spf13/cobra"
)

func Bitfinex(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	// Create a new bitfinex client
	c := bitfinex.NewClient(
		os.Getenv("BFX_API_KEY"),
		os.Getenv("BFX_API_SECRET"),
	)

	data, err := c.CallAPI(ctx, bitfinex.Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/ticker",
		SecType:  bitfinex.SecTypePrivate,
		Params:   map[string]interface{}{},
	}.Build())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
	// data, err := c.GetWallets(ctx)
	// data, err := c.GetActiveFundingOffers(ctx, "UST")
	// data, err := c.GetFundingStats(ctx, "fUSD")
	// data, err := c.GetWallets(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// b, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(b))
}
