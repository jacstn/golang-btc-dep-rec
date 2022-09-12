package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Deposit struct {
	TxId          string
	Vout          string
	Address       string
	Amount        float64
	Category      string
	Confirmations int8
}

type Customer struct {
	Id   uint64
	Name string
}

func NewCustomer(db *sql.DB, c *Customer) (int64, error) {
	res, err := db.Exec("INSERT INTO `customer` (name) values (?)", c.Name)

	if err != nil {
		fmt.Println("error while inserting into database", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		println("Error:", err.Error())
		return 0, err
	}

	return id, nil
}

func (c *Customer) Save(db *sql.DB, a *Address) error {
	var err error

	if c.Id == 0 {
		_, err = db.Exec("INSERT INTO `customer` (name) values (?)", a.CustomerId, a.Address)
		if err != nil {
			log.Println(err)
			return errors.New("Customer, unable to insert new record")
		}
		return nil
	}
	_, err = db.Exec("Update `customer` set name=?, address=?", a.CustomerId, a.Address)

	if err != nil {
		log.Println(err)
		return errors.New("Customer, unable to update record")
	}
	return nil
}

func ListCustomers(db *sql.DB) ([]Customer, error) {
	res, err := db.Query("SELECT * FROM `order` ORDER BY createdAt DESC LIMIT 20")
	if err != nil {
		fmt.Println("error while selecting orders from database")
		return []Customer{}, err
	}

	var customers []Customer

	for res.Next() {
		var c Customer
		err := res.Scan(&c.Id, &c.Name)

		if err != nil {
			log.Println(err)
			return customers, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

type Address struct {
	CustomerId string
	Address    string
}

func (a *Address) Save(db *sql.DB) error {
	_, err := db.Exec("INSERT into `address` set cucstomerId=?, address=?", a.CustomerId, a.Address)

	if err != nil {
		log.Println(err)
		return errors.New("Address, unable to insert record")
	}
	return nil
}
