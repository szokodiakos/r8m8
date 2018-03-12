package sql

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// DB interface
type DB interface {
	Begin() (Transaction, error)
}

type db struct {
	db *sqlx.DB
}

func (d *db) Begin() (Transaction, error) {
	log.Println("Transaction Begin")
	var transaction Transaction

	tx, err := d.db.Beginx()
	if err != nil {
		return transaction, err
	}

	transaction = NewSQLTransaction(tx)
	return transaction, nil
}

// NewSQLDB factory
func NewSQLDB(sqlDB *sqlx.DB) DB {
	return &db{
		db: sqlDB,
	}
}
