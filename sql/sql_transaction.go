package sql

import (
	"database/sql"

	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/szokodiakos/r8m8/logger"
)

// Transaction interface
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Commit() error
	Rollback() error
}

type transaction struct {
	tx *sqlx.Tx
}

func (t *transaction) Commit() error {
	logger.Get().Info("Transaction Commit")
	return t.tx.Commit()
}

func (t *transaction) Rollback() error {
	logger.Get().Info("Transaction Rollback")
	return t.tx.Rollback()
}

func (t *transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	logger.Get().Info(query, spew.Sdump(args))
	return t.tx.Exec(query, args...)
}

func (t *transaction) Select(dest interface{}, query string, args ...interface{}) error {
	logger.Get().Info(query, spew.Sdump(args))
	err := t.tx.Select(dest, query, args...)
	logger.Get().Info(spew.Sdump(dest))
	return err
}

func (t *transaction) Get(dest interface{}, query string, args ...interface{}) error {
	logger.Get().Info(query, spew.Sdump(args))
	err := t.tx.Get(dest, query, args...)
	logger.Get().Info(spew.Sdump(dest))
	return err
}

// NewSQLTransaction factory
func NewSQLTransaction(tx *sqlx.Tx) Transaction {
	return &transaction{
		tx: tx,
	}
}
