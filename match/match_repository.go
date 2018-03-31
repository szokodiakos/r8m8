package match

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Add(tr transaction.Transaction, match entity.Match) (entity.Match, error)
	GetByID(tr transaction.Transaction, matchID int64) (entity.Match, error)
	GetLatestByReporterPlayerID(tr transaction.Transaction, reporterPlayerID string) (entity.Match, error)
	Remove(tr transaction.Transaction, match entity.Match) error
}
