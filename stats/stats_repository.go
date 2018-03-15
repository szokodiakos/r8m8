package stats

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	GetLeaderboardPlayersByLeagueUniqueName(transaction transaction.Transaction, uniqueName string) ([]LeaderboardPlayer, error)
}
