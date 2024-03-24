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
	configPath  string
	mode        string
	searchFlag  string
	urlFlag     string
	parseURL    string
	dbTableName string
	pages       int
)

// ConfigFile
func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
	flag.StringVar(&mode, "mode", "web", "mode to run in. <web> makes HTTP requests, while <local> searches for .html file")
	flag.StringVar(&searchFlag, "search", "", "search")
	flag.StringVar(&urlFlag, "url", "", "parse url")
	flag.StringVar(&dbTableName, "table-name", "product_data", "database table name")
	flag.IntVar(&pages, "pages", 1, "how many pages to parse")
}
func main() {
	flag.Parse()
	config, isEnvLoaded := sparser.NewConfig()
	if !isEnvLoaded {
		_, err := toml.DecodeFS(cfg.ConfigFile, "config.toml", config)
		if err != nil {
			log.Fatal(err)
		}
	}
	config.DatabaseTableName = dbTableName
	// Setting the URL depending on input flags
	// If "-search" is passed -> search url
	// If "-url" is passed -> parsing specific catalog
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
		parseURL = "https://megamarket.ru/catalog/uhod-za-licom/"
	}

	// Send a GET request to the URL and parse
	if err := sparser.Start(config, parseURL, mode, pages); err != nil {
		log.Fatal(err)
		return
	}
}
