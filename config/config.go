package config

import (
	"log"
	"os"
)

type Config struct {
	ApiToken string
}

func LoadConfig() *Config {
	apiToken := os.Getenv("Pexels_Token")
	if apiToken == "" {
		log.Fatal("Pexels API Token is not set in the enviorment variables.")
	}

	return &Config{
		ApiToken: apiToken,
	}
}
