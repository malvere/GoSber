package model

import "database/sql"

type Product struct {
	ID           sql.NullInt64
	Title        string
	Price        int
	BonusAmount  int
	BonusPercent int
	Discount     int
	ProductID    int
	Link         string
}
