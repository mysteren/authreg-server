package apiserver

import (
	"os"

	"gitlab.devkeeper.com/authreg/server/internal/app/store"
)

type Config struct {
	BindAddr string
	LogLevel string
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: getEnv("BIND_ADDR", ":8080"),
		LogLevel: getEnv("LOG_LEVEL", "debug"),
		Store: store.NewConfig(
			getEnv("DB_CONNECT", ""),
			getEnv("DB_NAME", ""),
		),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
