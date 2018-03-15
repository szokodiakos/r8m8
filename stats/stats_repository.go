package stats

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	GetLeaderboardPlayersByLeagueUniqueName(tr transaction.Transaction, uniqueName string) ([]LeaderboardPlayer, error)
}
