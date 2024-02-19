package master

import (
	"github.com/dasbd72/asset-management/binance"
	"github.com/dasbd72/asset-management/okx"
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
