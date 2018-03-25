package entity

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// MatchRepository interface
type MatchRepository interface {
	Add(tr transaction.Transaction, match Match) (Match, error)
	GetByID(tr transaction.Transaction, matchID int64) (Match, error)
}
