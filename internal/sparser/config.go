package sparser

import "os"

type Config struct {
	DatabaseURL       string `toml:"database_url"`
	SQLiteURL         string `toml:"sqlite_url"`
	DatabaseTableName string
}

func NewConfig() (*Config, bool) {
	url := os.Getenv("DB_URL")
	if url != "" {
		return &Config{DatabaseURL: url}, true
	} else {
		return &Config{}, false
	}

}

func LoadEnv() *Config {
	url := os.Getenv("DB_URL")
	return &Config{
		DatabaseURL: url,
	}
}
