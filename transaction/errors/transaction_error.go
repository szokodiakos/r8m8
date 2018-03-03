package errors

import (
	"fmt"
)

type transactionError struct {
	err error
}

func (e *transactionError) Error() string {
	return fmt.Sprintf("Transaction Error: %s", e.err)
}

// NewTransactionError factory
func NewTransactionError(err error) error {
	return &transactionError{
		err: err,
	}
}
