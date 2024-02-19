package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/dasbd72/asset-management/okx"
	"github.com/spf13/cobra"
)

func OKX(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	// Create a new OKX client
	c := okx.NewClient(
		os.Getenv("OKX_API_KEY"),
		os.Getenv("OKX_API_SECRET"),
		os.Getenv("OKX_PASSPHRASE"),
	)

	// Start testing
	// data, err := c.GetBalance(ctx, &okx.GetBalanceRequest{
	// 	Ccy: "BTC",
	// })
	// data, err := c.GetTicker(ctx, &okx.GetTickerRequest{
	// 	InstID: "BTC-USDT",
	// })
	// data, err := c.GetTickers(ctx, &okx.GetTickersRequest{
	// 	InstType: "SPOT",
	// })
	// data, err := c.GetFundingBalances(ctx, &okx.GetFundingBalancesRequest{})
	data, err := c.GetSavingBalance(ctx, &okx.GetSavingBalanceRequest{})
	// data, err := c.GetETHStakingBalance(ctx, &okx.GetETHStakingBalanceRequest{})
	// data, err := c.GetEarnOffers(ctx, &okx.GetEarnOffersRequest{})
	// data, err := c.GetActiveEarnOrders(ctx, &okx.GetActiveEarnOrdersRequest{})
	// data, err := c.GetLendingHistory(ctx, &okx.GetLendingHistoryRequest{})
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
