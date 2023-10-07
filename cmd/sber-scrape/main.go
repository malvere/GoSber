package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/url"
	cfg "sber-scrape/config"
	"sber-scrape/internal/sparser"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

// ConfigFile
func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := sparser.NewConfig()
	_, err := toml.DecodeFS(cfg.ConfigFile, "config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	// Define the URL
	tag := url.QueryEscape("Моторное масло 4л")
	url := fmt.Sprintf("https://megamarket.ru/catalog/?q=%s", tag)

	// Send a GET request to the URL

	if err := sparser.Start(config, url); err != nil {
		log.Fatal(err)
	}
}
