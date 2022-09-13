package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Customer struct {
	Name    string
	Address string
}

func (c *Customer) Save(db *sql.DB) error {
	var err error

	err = db.QueryRow("SELECT * from customer where address=??", c.Address).Scan(&c.Address)

	if err != nil {
		// record already exists, update confirmations only
		_, err = db.Exec("INSERT INTO `customer` (name, address) values (?, ?)", c.Name, c.Address)
		if err != nil {
			log.Println(err)
			return errors.New("Customer, unable to insert new record")
		}
		return nil
	}

	_, err = db.Exec("Update `customer` set name=?, address=?", c.Name, c.Address)

	if err != nil {
		log.Println(err)
		return errors.New("Customer, unable to update record")
	}
	return nil
}

func ListCustomers(db *sql.DB) ([]Customer, error) {
	res, err := db.Query("SELECT address, name FROM `customer`")
	if err != nil {
		fmt.Println("error while selecting customers from database")
		return []Customer{}, err
	}

	var customers []Customer

	for res.Next() {
		var c Customer
		err := res.Scan(&c.Address, &c.Name)

		if err != nil {
			log.Println(err)
			return customers, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

type Deposit struct {
	TxId          string
	Vout          int16
	Address       string
	Amount        float64
	Category      string
	Confirmations int16
}
