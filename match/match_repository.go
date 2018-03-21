package match

import (
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, match model.Match) (int64, error)
	GetByID(tr transaction.Transaction, matchID int64) (model.Match, error)
}
