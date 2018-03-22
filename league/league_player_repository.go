package league

import (
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerRepository interface
type PlayerRepository interface {
	GetMultipleByPlayerUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]model.LeaguePlayer, error)
	GetMultipleByLeagueUniqueNameOrderedByRating(tr transaction.Transaction, uniqueName string) ([]model.LeaguePlayer, error)
	Update(tr transaction.Transaction, leaguePlayer model.LeaguePlayer) error
	Create(tr transaction.Transaction, leaguePlayer model.LeaguePlayer) error
}
