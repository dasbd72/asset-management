package main

import (
	"encoding/json"
	"log"

	"github.com/dasbd72/go-exchange-sdk/max"
	"github.com/spf13/cobra"
)

func Max(cmd *cobra.Command, args []string) {
	// Start testing
	data, err := max.GetUsdtToTWD()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
