package sparser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sber-scrape/internal/model"
	"sber-scrape/internal/store"
)

func createFile(fileName string) (*os.File, *csv.Writer, error) {
	headers := []string{
		"Title",
		"Price",
		"BonusAmount",
		"BonusPercent",
		"Discount",
		"ProductID",
		"Link",
	}
	file, err := os.Create(fmt.Sprintf("%s.csv", fileName))
	if err != nil {
		log.Print(err)
	}

	writer := csv.NewWriter(file)
	err = writer.Write(headers)
	if err != nil {
		log.Print(err)
		return nil, nil, err
	}
	return file, writer, nil
}

func CommaSeparated(fileName string, store store.Store) error {
	file, writer, err := createFile("sber")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer writer.Flush()

	p := &model.Product{}
	rows, err := store.Product().FindAll()
	if err != nil {
		log.Fatal(err)
		return err
	}
	for rows.Next() {
		err := rows.Scan(
			&p.Title,
			&p.Price,
			&p.BonusAmount,
			&p.BonusPercent,
			&p.Discount,
			&p.ProductID,
			&p.Link,
		)
		if err != nil {
			log.Fatal(err)
			return err
		}

		row := []string{
			fmt.Sprint(p.Title),
			fmt.Sprint(p.Price),
			fmt.Sprint(p.BonusAmount),
			fmt.Sprint(p.BonusPercent),
			fmt.Sprint(p.Discount),
			fmt.Sprint(p.ProductID),
			fmt.Sprint(p.Link),
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}
	return nil
}
