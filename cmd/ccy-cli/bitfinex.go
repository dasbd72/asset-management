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

	res, err := c.CallAPI(ctx, bitfinex.Request_builder{
		Method:   http.MethodPost,
		Endpoint: "/auth/r/wallets",
		SecType:  bitfinex.SecTypePrivate,
		Params:   map[string]interface{}{},
	}.Build())
	// res, err := c.CallAPI(ctx, bitfinex.Request_builder{
	// 	Method:      http.MethodGet,
	// 	Endpoint:    "/ticker",
	// 	SubEndpoint: "/tBTCUSD",
	// 	SecType:     bitfinex.SecTypePrivate,
	// 	Params:      map[string]interface{}{},
	// }.Build())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(res))
}
