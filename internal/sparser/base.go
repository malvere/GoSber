package sparser

import (
	"database/sql"
	"errors"
	"log"
	"sber-scrape/internal/store"
	"sber-scrape/internal/store/litestore"
	"sber-scrape/internal/store/sqlstore"
)

func Start(config *Config, url string, mode string) error {
	store, db, err := selDB(config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	switch mode {
	case "web":
		if err := GetPages(url, store, 6); err != nil {
			log.Fatal(err)
		}
	case "local":
		if err := GetLocalHtml("page.html", store); err != nil {
			log.Fatal(err)
		}
	}

	if store, ok := store.(*litestore.Store); ok {
		log.Print("Preparing .csv file...")
		CommaSeparated("sber", store)
		log.Print("File Written!")
	} else {
		log.Print("No .csv file will be created: PostgreSQL is active")
	}
	return nil
}

func selDB(config *Config) (store.Store, *sql.DB, error) {
	var s store.Store
	databaseDrivers := map[string]string{
		"postgres": config.DatabaseURL,
		"sqlite":   config.SQLiteURL,
	}
	for driver, url := range databaseDrivers {
		db, err := sql.Open(driver, url)
		if err != nil {
			log.Printf("Driver '%s' is not active, using", driver)
			continue
		}
		if err := db.Ping(); err != nil {
			log.Printf("Could not ping '%s' driver", driver)
			continue
		}

		if driver == "postgres" {
			s = sqlstore.New(db)
		} else {
			s = litestore.New(db)
		}

		err = s.Product().NewTable()
		if err != nil {
			log.Printf("Something went wrong while creating tables: %s", err)
		}
		log.Printf("Driver '%s' is active", driver)
		return s, db, nil
	}
	return nil, nil, errors.New("could not connect to DB")
}
