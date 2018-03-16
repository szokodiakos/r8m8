package stats

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	GetPlayersStatsByLeagueUniqueName(tr transaction.Transaction, uniqueName string) ([]PlayerStats, error)
}
