package config

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Config struct {
	BinanceApiKey    string
	BinanceApiSecret string
}

func Load() *Config {
	godotenv.Load()
	godotenv.Load(path.Join(os.Getenv("HOME"), ".ccyrc"))

	return &Config{
		BinanceApiKey:    os.Getenv("BINANCE_API_KEY"),
		BinanceApiSecret: os.Getenv("BINANCE_API_SECRET"),
	}
}
