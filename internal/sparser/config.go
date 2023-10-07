package sparser

type Config struct {
	DatabaseURL string `toml:"database_url"`
	SQLiteURL   string `toml:"sqlite_url"`
}

func NewConfig() *Config {
	return &Config{}
}
