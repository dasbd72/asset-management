package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	binanceSpot "github.com/dasbd72/go-exchange-sdk/binance/pkg/spot"
	"github.com/spf13/cobra"
)

func Binance(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	// Create a new Binance client
	c := binanceSpot.NewClient(
		os.Getenv("BINANCE_API_KEY"),
		os.Getenv("BINANCE_API_SECRET"),
	)

	// Start testing
	data, err := c.GetUserWalletBalance(ctx)
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
