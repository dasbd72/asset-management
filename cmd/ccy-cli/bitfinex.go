package main

import (
	"context"
	"encoding/json"
	"log"
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

	// res, err := c.CallAPI(ctx, bitfinex.Request_builder{
	// 	Method:      http.MethodGet,
	// 	Endpoint:    "/ticker",
	// 	SubEndpoint: "/tBTCUSD",
	// 	SecType:     bitfinex.SecTypePrivate,
	// 	Params:      map[string]interface{}{},
	// }.Build())
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
