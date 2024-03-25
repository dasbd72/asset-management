package manager

import (
	binanceRest "github.com/dasbd72/go-exchange-sdk/binance/pkg/spot"
	bitfinexRest "github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/rest"
	okxRest "github.com/dasbd72/go-exchange-sdk/okx/pkg/rest"
)

type Client struct {
	binanceClient      *binanceRest.Client
	okxClient          *okxRest.Client
	bitfinexRestClient *bitfinexRest.Client
}

type Client_builder struct {
	BinanceClient      *binanceRest.Client
	OkxClient          *okxRest.Client
	BitfinexRestClient *bitfinexRest.Client
}

func (b Client_builder) Build() *Client {
	return &Client{
		binanceClient:      b.BinanceClient,
		okxClient:          b.OkxClient,
		bitfinexRestClient: b.BitfinexRestClient,
	}
}
