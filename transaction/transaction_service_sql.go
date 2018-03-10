package transaction

import (
	"github.com/szokodiakos/r8m8/sql"

	"github.com/szokodiakos/r8m8/transaction/errors"
)

type transactionService struct {
	db sql.DB
}

func (ts *transactionService) Start() (Transaction, error) {
	var tr Transaction
	tx, err := ts.db.Begin()
	if err != nil {
		return tr, errors.NewTransactionError(err)
	}
	tr = Transaction{
		transaction: tx,
	}
	return tr, nil
}

func (ts *transactionService) Commit(tr Transaction) error {
	tx := tr.transaction.(sql.Transaction)
	err := tx.Commit()
	if err != nil {
		return errors.NewTransactionError(err)
	}
	return nil
}

func (ts *transactionService) Rollback(tr Transaction) error {
	tx := tr.transaction.(sql.Transaction)
	err := tx.Rollback()
	if err != nil {
		return errors.NewTransactionError(err)
	}
	return nil
}

// NewServiceSQL factory
func NewServiceSQL(db sql.DB) Service {
	return &transactionService{
		db: db,
	}
}
