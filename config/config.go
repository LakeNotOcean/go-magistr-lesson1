package config

import "os"

type Config struct {
	MetricUrl string
}

func NewConfig() *Config {
	return &Config{
		MetricUrl: getEnv("METRIC_URL", "srv.msk01.gigacorp.local"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
