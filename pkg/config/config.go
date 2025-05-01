package config

import (
	"os"
)

type Config struct {
	APIKey     string
	BaseURL    string
	Symbol     string
	OutputSize string
	DataType   string
	WindowSize int
}

func NewConfig() *Config {
	return &Config{
		APIKey:     getEnvOrDefault("ALPHA_VANTAGE_API_KEY", "YOUR_API_KEY"),
		BaseURL:    "https://www.alphavantage.co/query",
		Symbol:     "MSFT",
		OutputSize: "compact",
		DataType:   "json",
		WindowSize: 5,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 