package apiserver

import (
	"os"
)

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: getEnv("BIND_ADDR", ":8080"),
		LogLevel: getEnv("LOG_LEVEL", "debug"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
