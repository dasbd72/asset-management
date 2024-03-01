package manager

import (
	"github.com/dasbd72/go-exchange-sdk/binance"
	"github.com/dasbd72/go-exchange-sdk/okx"
)

type Client struct {
	binanceClient *binance.Client
	okxClient     *okx.Client
}

func NewClient(binanceClient *binance.Client, okxClient *okx.Client) *Client {
	return &Client{
		binanceClient: binanceClient,
		okxClient:     okxClient,
	}
}
