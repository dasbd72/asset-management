module github.com/dasbd72/go-exchange-sdk/cmd/ccy-cli

go 1.22.0

replace github.com/dasbd72/go-exchange-sdk/binance => ../../binance

replace github.com/dasbd72/go-exchange-sdk/okx => ../../okx

replace github.com/dasbd72/go-exchange-sdk/max => ../../max

replace github.com/dasbd72/go-exchange-sdk/bitfinex => ../../bitfinex

replace github.com/dasbd72/go-exchange-sdk/manager => ../../manager

require (
	github.com/dasbd72/go-exchange-sdk/binance v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/bitfinex v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/manager v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/max v0.0.0-00010101000000-000000000000
	github.com/dasbd72/go-exchange-sdk/okx v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/spf13/cobra v1.8.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
