package config

import "os"

type Config struct {
	AppPort string
	AppEnv  string
}

func Load() Config {
	return Config{
		AppPort: getEnv("APP_PORT", "8081"),
		AppEnv:  getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
