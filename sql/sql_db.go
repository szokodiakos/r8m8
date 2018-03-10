package sql

import (
	"database/sql"
	"log"
)

// DB interface
type DB interface {
	Begin() (Transaction, error)
}

type db struct {
	db *sql.DB
}

func (d *db) Begin() (Transaction, error) {
	log.Println("Transaction Begin")
	var transaction Transaction

	tx, err := d.db.Begin()
	if err != nil {
		return transaction, err
	}

	transaction = NewSQLTransaction(tx)
	return transaction, nil
}

// NewSQLDB factory
func NewSQLDB(sqlDB *sql.DB) DB {
	return &db{
		db: sqlDB,
	}
}
