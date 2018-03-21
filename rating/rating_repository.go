package rating

import (
	"github.com/szokodiakos/r8m8/rating/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, rating model.Rating) error
	GetMultipleByPlayerIDs(tr transaction.Transaction, playerIDs []int64) ([]model.Rating, error)
	Update(tr transaction.Transaction, rating model.Rating) error
}
