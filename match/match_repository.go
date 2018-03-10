package match

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	Create(transaction transaction.Transaction) (int64, error)
}
