package manager

import (
	"github.com/dasbd72/go-exchange-sdk/binance"
	"github.com/dasbd72/go-exchange-sdk/bitfinex"
	"github.com/dasbd72/go-exchange-sdk/okx"
)

type Client struct {
	binanceClient  *binance.Client
	okxClient      *okx.Client
	bitfinexClient *bitfinex.Client
}

type Client_builder struct {
	BinanceClient  *binance.Client
	OkxClient      *okx.Client
	BitfinexClient *bitfinex.Client
}

func (b Client_builder) Build() *Client {
	return &Client{
		binanceClient:  b.BinanceClient,
		okxClient:      b.OkxClient,
		bitfinexClient: b.BitfinexClient,
	}
}
