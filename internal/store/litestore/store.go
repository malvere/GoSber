package litestore

import (
	"database/sql"
	"sber-scrape/internal/store"

	_ "modernc.org/sqlite"
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
