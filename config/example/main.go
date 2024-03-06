package main

import (
	"encoding/json"
	"fmt"

	"github.com/dasbd72/go-exchange-sdk/config"
)

func main() {
	cfg := config.Load()
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
