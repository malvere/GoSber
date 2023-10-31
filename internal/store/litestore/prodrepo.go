package litestore

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"sber-scrape/internal/model"
	"sber-scrape/internal/store"

	"github.com/lib/pq"
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
	rows, err := r.store.db.Query(
		"SELECT title, price, bonuses, bonus_percent, discount, product_id, link FROM product_data",
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rows, nil
}

func (r *ProductRepo) NewTable() error {
	_, err := r.store.db.Exec(
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

func (r *ProductRepo) BulkInsertProducts(products []model.Product) error {
	// Create a buffer to store the CSV data
	var csvData bytes.Buffer

	// Write the CSV header
	csvData.WriteString("title,price,bonuses,bonus_percent,discount,product_id,link\n")

	// Batch size
	batchSize := 100
	productCount := len(products)

	for i := 0; i < productCount; i += batchSize {
		batchEnd := i + batchSize
		if batchEnd > productCount {
			batchEnd = productCount
		}

		// Write the data for each product in the batch
		for _, p := range products[i:batchEnd] {
			csvData.WriteString(fmt.Sprintf("%s,%d,%d,%d,%d,%d,%s\n", p.Title, p.Price, p.BonusAmount, p.BonusPercent, p.Discount, p.ProductID, p.Link))
		}

		// Use the COPY command to bulk insert data for this batch
		_, err := r.store.db.Exec(`COPY product_data FROM STDIN WITH CSV HEADER`, csvData.String())
		if err != nil {
			return err
		}

		// Reset the buffer for the next batch
		csvData.Reset()
	}

	return nil
}

func (r *ProductRepo) Transaction(products []model.Product) error {
	txn, err := r.store.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := txn.Prepare(pq.CopyIn(
		"product_data",
		"title",
		"price",
		"bonuses",
		"bonus_percent",
		"discount",
		"product_id",
		"link",
	))
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range products {
		_, err := stmt.Exec(p.Title, p.Price, p.BonusAmount, p.BonusPercent, p.Discount, p.ProductID, p.Link)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil

}
