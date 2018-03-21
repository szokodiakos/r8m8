package league

import (
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerRepository interface
type PlayerRepository interface {
	GetMultipleByLeagueUniqueNameOrderedByRating(tr transaction.Transaction, uniqueName string) ([]model.LeaguePlayer, error)
}
