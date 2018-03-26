package transaction

import (
	"github.com/szokodiakos/r8m8/sql"
)

type transactionServiceSQL struct {
	db sql.DB
}

func (ts *transactionServiceSQL) Start() (Transaction, error) {
	var transaction Transaction
	sqlTransaction, err := ts.db.Begin()
	if err != nil {
		return transaction, err
	}
	transaction = Transaction{
		concreteTransaction: sqlTransaction,
	}
	return transaction, nil
}

func (ts *transactionServiceSQL) Commit(transaction Transaction) error {
	sqlTransaction := GetSQLTransaction(transaction)
	return sqlTransaction.Commit()
}

func (ts *transactionServiceSQL) Rollback(transaction Transaction, err error) error {
	sqlTransaction := GetSQLTransaction(transaction)
	trErr := sqlTransaction.Rollback()
	if trErr != nil {
		return trErr
	}
	return err
}

// NewServiceSQL factory
func NewServiceSQL(db sql.DB) Service {
	return &transactionServiceSQL{
		db: db,
	}
}
