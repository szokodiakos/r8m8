package rating

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Create(transaction transaction.Transaction, rating Rating) error
	GetMultipleByPlayerIDs(transaction transaction.Transaction, playerIDs []int64) ([]RepoRating, error)
	UpdateRating(transaction transaction.Transaction, rating Rating) error
}
