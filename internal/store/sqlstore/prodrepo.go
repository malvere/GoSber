package sqlstore

import (
	"context"
	"database/sql"
	"log"
	"sber-scrape/internal/model"
	"sber-scrape/internal/store"
)

type ProductRepo struct {
	store *Store
}

func (r *ProductRepo) Create(p *model.Product) error {
	return r.store.db.QueryRow(
		"INSERT INTO product_data (title, price, bonuses, bonus_percent, discount, product_id, link) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		p.Title,
		p.Price,
		p.BonusAmount,
		p.BonusPercent,
		p.Discount,
		p.ProductID,
		p.Link,
	).Scan(&p.ID)
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
