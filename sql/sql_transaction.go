package sql

import (
	"database/sql"
	"log"
)

// Transaction interface
type Transaction interface {
	Commit() error
	Rollback() error
}

type transaction struct {
	tx *sql.Tx
}

func (t *transaction) Commit() error {
	log.Println("Transaction Commit")
	return t.tx.Commit()
}

func (t *transaction) Rollback() error {
	log.Println("Transaction Rollback")
	return t.tx.Rollback()
}

// NewSQLTransaction factory
func NewSQLTransaction(tx *sql.Tx) Transaction {
	return &transaction{
		tx: tx,
	}
}
