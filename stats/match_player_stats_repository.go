package stats

import (
	"github.com/szokodiakos/r8m8/stats/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// MatchPlayerStatsRepository interface
type MatchPlayerStatsRepository interface {
	GetMultipleByMatchID(tr transaction.Transaction, matchID int64) ([]model.MatchPlayerStats, error)
}
