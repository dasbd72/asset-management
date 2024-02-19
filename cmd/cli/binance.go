package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/dasbd72/asset-management/binance"
	"github.com/spf13/cobra"
)

func Binance(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	// Create a new Binance client
	c := binance.NewClient(
		os.Getenv("BINANCE_API_KEY"),
		os.Getenv("BINANCE_API_SECRET"),
	)

	// Start testing
	data, err := c.GetOrderBook(ctx, binance.NewGetOrderBookRequest("BTCUSDT"))
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(data)
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
