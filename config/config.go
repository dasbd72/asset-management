package config

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Config struct {
	BinanceApiKey     string
	BinanceApiSecret  string
	OKXApiKey         string
	OKXApiSecret      string
	OKXPassphrase     string
	BitfinexApiKey    string
	BitfinexApiSecret string
	PionexApiKey      string
	PionexApiSecret   string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	err = godotenv.Load(path.Join(os.Getenv("HOME"), ".ccyrc"))
	if err != nil {
		return nil, err
	}

	return &Config{
		BinanceApiKey:     os.Getenv("BINANCE_API_KEY"),
		BinanceApiSecret:  os.Getenv("BINANCE_API_SECRET"),
		OKXApiKey:         os.Getenv("OKX_API_KEY"),
		OKXApiSecret:      os.Getenv("OKX_API_SECRET"),
		OKXPassphrase:     os.Getenv("OKX_PASSPHRASE"),
		BitfinexApiKey:    os.Getenv("BFX_API_KEY"),
		BitfinexApiSecret: os.Getenv("BFX_API_SECRET"),
		PionexApiKey:      os.Getenv("PIONEX_API_KEY"),
		PionexApiSecret:   os.Getenv("PIONEX_API_SECRET"),
	}, nil
}
