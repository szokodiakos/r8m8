package player

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	Create(transaction transaction.Transaction) (int64, error)
	UpdateRatingByID(transaction transaction.Transaction, ID int64, rating int) error
}
