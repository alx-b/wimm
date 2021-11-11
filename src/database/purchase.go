package database

import (
	"fmt"
)

type PurchaseInDB struct {
	Name   string
	Seller string
	Tag    string
	Cost   float64
	Date   string
}

type PurchaseOutDB struct {
	Id     int
	Name   string
	Seller string
	Tag    string
	Cost   float64
	Date   string
}

// Add Purchase to database
func (d *Database) AddPurchase(purchase PurchaseInDB) error {
	_, err := d.conn.Exec(
		"INSERT INTO purchase (name, seller, tag, cost, date) VALUES (?,?,?,?,?)",
		purchase.Name,
		purchase.Seller,
		purchase.Tag,
		purchase.Cost,
		purchase.Date,
	)

	if err != nil {
		return fmt.Errorf("%s: %s", ErrCouldNotQueryDatabase, err.Error())
	}

	return nil
}

// Get all purchases from database
func (d *Database) GetPurchases() ([]PurchaseOutDB, error) {
	rows, err := d.conn.Query("SELECT * FROM purchase")

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrCouldNotQueryDatabase, err.Error())
	}

	list := []PurchaseOutDB{}

	for rows.Next() {
		purchase := PurchaseOutDB{}
		rows.Scan(
			&purchase.Id,
			&purchase.Name,
			&purchase.Seller,
			&purchase.Tag,
			&purchase.Cost,
			&purchase.Date,
		)
		list = append(list, purchase)
	}
	return list, nil
}
