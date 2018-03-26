package sql

import (
	"database/sql"

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
	logger.Get().Info("Transaction Begin")
	var transaction Transaction

	tx, err := d.db.Beginx()
	if err != nil {
		return transaction, err
	}

	transaction = NewSQLTransaction(tx)
	return transaction, nil
}

// NewSQLDB factory
func NewSQLDB(sqlDB *sql.DB, dialect string) DB {
	return &db{
		db: sqlx.NewDb(sqlDB, dialect),
	}
}
