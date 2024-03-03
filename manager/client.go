package manager

import (
	binanceRest "github.com/dasbd72/go-exchange-sdk/binance/pkg/spot"
	bitfinexRest "github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/rest"
	"github.com/dasbd72/go-exchange-sdk/okx"
)

type Client struct {
	binanceClient      *binanceRest.Client
	okxClient          *okx.Client
	bitfinexRestClient *bitfinexRest.Client
}

type Client_builder struct {
	BinanceClient      *binanceRest.Client
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
