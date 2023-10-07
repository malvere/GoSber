package sqlstore

import (
	"database/sql"
	"log"
	"sber-scrape/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db          *sql.DB
	productRepo *ProductRepo
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Product() store.ProductRepo {
	if s.productRepo != nil {
		return s.productRepo
	}
	s.productRepo = &ProductRepo{
		store: s,
	}
	return s.productRepo

}

func (r *ProductRepo) NewTable() error {
	_, err := r.store.db.Query(
		`CREATE TABLE IF NOT EXISTS product_data (
			id SERIAL PRIMARY KEY,
			title TEXT,
			price INT,
			bonuses INT,
			bonus_percent INT,
			discount INT,
			product_id BIGINT,
			link TEXT
		)`,
	)
	if err != nil {
		log.Fatalln("Error creating SQLite Table: ", err)
		return err
	}
	return nil
}
