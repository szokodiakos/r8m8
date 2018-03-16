package rating

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, rating Rating) error
	GetMultipleByPlayerIDs(tr transaction.Transaction, playerIDs []int64) ([]Rating, error)
	UpdateRating(tr transaction.Transaction, rating Rating) error
}
