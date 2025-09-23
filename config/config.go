package config

import "os"

type Config struct {
	MetricUrl string
	Scheme    string
}

func NewConfig() *Config {
	return &Config{
		MetricUrl: getEnv("METRIC_URI", "srv.msk01.gigacorp.local"),
		Scheme:    getEnv("SCHEME", "http"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
