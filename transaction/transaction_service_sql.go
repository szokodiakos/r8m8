package transaction

import (
	"database/sql"

	"github.com/szokodiakos/r8m8/transaction/errors"
)

type transactionService struct {
	db *sql.DB
}

func (ts *transactionService) Start() (Transaction, error) {
	var tr Transaction
	tx, err := ts.db.Begin()
	if err != nil {
		return tr, errors.NewTransactionError(err)
	}
	tr = Transaction{
		transaction: &tx,
	}
	return tr, nil
}

func (ts *transactionService) Commit(tr Transaction) error {
	tx := tr.transaction.(*sql.Tx)
	err := tx.Commit()
	if err != nil {
		return errors.NewTransactionError(err)
	}
	return nil
}

func (ts *transactionService) Rollback(tr Transaction) error {
	tx := tr.transaction.(*sql.Tx)
	err := tx.Rollback()
	if err != nil {
		return errors.NewTransactionError(err)
	}
	return nil
}

// NewService factory
func NewService(db *sql.DB) Service {
	return &transactionService{
		db: db,
	}
}
