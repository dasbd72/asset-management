module github.com/dasbd72/go-exchange-sdk/manager

go 1.22.0

replace github.com/dasbd72/go-exchange-sdk/binance => ../binance

replace github.com/dasbd72/go-exchange-sdk/okx => ../okx

replace github.com/dasbd72/go-exchange-sdk/bitfinex => ../bitfinex

replace github.com/dasbd72/go-exchange-sdk/max => ../max

require (
	github.com/dasbd72/go-exchange-sdk/binance v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/bitfinex v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/max v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/okx v0.0.0-00010101000000-000000000000
)
