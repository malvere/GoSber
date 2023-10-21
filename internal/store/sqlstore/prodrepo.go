package sqlstore

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sber-scrape/internal/model"
	"sber-scrape/internal/store"
)

type ProductRepo struct {
	store *Store
}

func (r *ProductRepo) Create(p *model.Product) error {
	_, err := r.store.db.Exec(
		"INSERT INTO product_data (title, price, bonuses, bonus_percent, discount, product_id, link) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		p.Title,
		p.Price,
		p.BonusAmount,
		p.BonusPercent,
		p.Discount,
		p.ProductID,
		p.Link,
	)
	// .Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) FindByProductId(productID int) (*model.Product, error) {
	p := &model.Product{}
	if err := r.store.db.QueryRow(
		"SELECT title, price, bonuses, bonus_percent, discount, product_id, link FROM product_data WHERE product_id = $1",
		productID,
	).Scan(
		&p.Title,
		&p.Price,
		&p.BonusAmount,
		&p.BonusPercent,
		&p.Discount,
		&p.ProductID,
		&p.Link,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecNotFound
		}
		return nil, err
	}
	return p, nil
}

func (r *ProductRepo) FindAll() (*sql.Rows, error) {
	ctx := context.Background()
	rows, err := r.store.db.QueryContext(ctx,
		"SELECT title, price, bonuses, bonus_percent, discount, product_id, link FROM product_data",
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rows, nil
}

func (r *ProductRepo) NewTable() error {
	// Check if the table exists
	query := "SELECT to_regclass('product_data')"
	var tableName string
	if err := r.store.db.QueryRow(query).Scan(&tableName); err != nil {
		fmt.Print("Found nothind, creating table >>>")
		// If the table doesn't exist, create it
		return r.createTable()
	}
	return nil
}

func (r *ProductRepo) createTable() error {
	_, err := r.store.db.Exec(`
			CREATE TABLE IF NOT EXISTS product_data (
				id SERIAL PRIMARY KEY,
				title TEXT,
				price INT,
				bonuses INT,
				bonus_percent INT,
				discount INT,
				product_id BIGINT,
				link TEXT
			)
		`)
	if err != nil {
		return err
	}
	return nil
}
