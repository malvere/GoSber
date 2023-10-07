package store

import (
	"database/sql"
	"sber-scrape/internal/model"
)

type ProductRepo interface {
	Create(*model.Product) error
	FindByProductId(int) (*model.Product, error)
	FindAll() (*sql.Rows, error)
	NewTable() error
}
