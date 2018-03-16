package stats

import "github.com/szokodiakos/r8m8/transaction"

// PlayerRepository interface
type PlayerRepository interface {
	GetMultipleByLeagueUniqueName(tr transaction.Transaction, uniqueName string) ([]PlayerStats, error)
}
