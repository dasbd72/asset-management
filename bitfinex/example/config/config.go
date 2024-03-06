package config

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Config struct {
	BitfinexApiKey    string
	BitfinexApiSecret string
}

func Load() *Config {
	godotenv.Load()
	godotenv.Load(path.Join(os.Getenv("HOME"), ".ccyrc"))

	return &Config{
		BitfinexApiKey:    os.Getenv("BFX_API_KEY"),
		BitfinexApiSecret: os.Getenv("BFX_API_SECRET"),
	}
}
