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
	mode       string
	searchFlag string
	urlFlag    string
	parseURL   string
)

// ConfigFile
func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
	flag.StringVar(&mode, "m", "web", "mode to run in. <web> makes HTTP requests, while <local> searches for .html file")
	flag.StringVar(&searchFlag, "s", "", "search")
	flag.StringVar(&urlFlag, "u", "", "parse url")
}
func main() {
	flag.Parse()

	config := sparser.NewConfig()
	_, err := toml.DecodeFS(cfg.ConfigFile, "config.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	// Setting the URL depending on input flags
	// If "-s" is passed -> search url
	// If "-u" is passed -> parsing specific catalog
	if searchFlag != "" {
		log.Println("Searching for: ", searchFlag)
		parseURL = fmt.Sprintf(
			"https://megamarket.ru/catalog/?q=%s",
			url.QueryEscape(searchFlag),
		)
	} else if urlFlag != "" {
		log.Println("Scraping: ", urlFlag)
		parseURL = urlFlag
	} else {
		parseURL = "https://megamarket.ru/catalog/?q=нарзан"
	}

	// Send a GET request to the URL and parse
	if err := sparser.Start(config, parseURL, mode); err != nil {
		log.Fatal(err)
	}
}
