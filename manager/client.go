package manager

import (
	"github.com/dasbd72/go-exchange-sdk/binance"
	bitfinexRest "github.com/dasbd72/go-exchange-sdk/bitfinex/rest"
	"github.com/dasbd72/go-exchange-sdk/okx"
)

type Client struct {
	binanceClient      *binance.Client
	okxClient          *okx.Client
	bitfinexRestClient *bitfinexRest.Client
}

type Client_builder struct {
	BinanceClient      *binance.Client
	OkxClient          *okx.Client
	BitfinexRestClient *bitfinexRest.Client
}

func (b Client_builder) Build() *Client {
	return &Client{
		binanceClient:      b.BinanceClient,
		okxClient:          b.OkxClient,
		bitfinexRestClient: b.BitfinexRestClient,
	}
}
