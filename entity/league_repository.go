package entity

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// LeagueRepository interface
type LeagueRepository interface {
	GetByID(tr transaction.Transaction, id string) (League, error)
	Add(tr transaction.Transaction, league League) (League, error)
	Update(tr transaction.Transaction, league League) error
}
