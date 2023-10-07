package sparser

import (
	"database/sql"
	"log"
	"sber-scrape/internal/store"
	"sber-scrape/internal/store/litestore"
	"sber-scrape/internal/store/sqlstore"
)

func Start(config *Config, url string) error {
	store, db, err := chooseStore(config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := GetHtml(url, store); err != nil {
		log.Fatal(err)
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

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("PG DB")
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Print("PING")
		log.Print(err)
		return nil, err
	}

	return db, nil
}

func newLiteDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", databaseURL)
	if err != nil {
		log.Fatal("SQ")
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Print("SQLite")
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func chooseStore(config *Config) (store.Store, *sql.DB, error) {
	// Try to connect to PostgreSQL
	postgresDB, err := newDB(config.DatabaseURL)
	if err == nil {
		log.Print("Postgres active")
		return sqlstore.New(postgresDB), postgresDB, nil
	}

	// If PostgreSQL is not available, use SQLite
	sqliteDB, err := newLiteDB(config.SQLiteURL)
	if err != nil {
		log.Fatal("SQLite not active")
		return nil, nil, err
	}
	s := litestore.New(sqliteDB)
	s.Product().NewTable()
	log.Print("SQLite active")
	return s, sqliteDB, nil
}
