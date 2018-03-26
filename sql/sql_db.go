package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/szokodiakos/r8m8/logger"
)

// DB interface
type DB interface {
	Begin() (Transaction, error)
}

type db struct {
	db *sqlx.DB
}

func (d *db) Begin() (Transaction, error) {
	logger.Get().Debug("Transaction Begin")
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
