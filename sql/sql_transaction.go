package sql

import (
	"database/sql"
	"log"
)

// Transaction interface
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
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

func (t *transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Println(query, args)
	return t.tx.Exec(query, args...)
}

func (t *transaction) Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Println(query, args)
	return t.tx.Query(query, args...)
}

func (t *transaction) QueryRow(query string, args ...interface{}) *sql.Row {
	log.Println(query, args)
	return t.tx.QueryRow(query, args...)
}

// NewSQLTransaction factory
func NewSQLTransaction(tx *sql.Tx) Transaction {
	return &transaction{
		tx: tx,
	}
}
