package entity

import (
	"sort"

	"github.com/szokodiakos/r8m8/league/errors"
)

// LeagueStats struct
type LeagueStats struct {
	LeaguePlayer LeaguePlayer
	Place        int
}

// LeaguePlayers type
type LeaguePlayers []LeaguePlayer

// League struct
type League struct {
	ID            string `db:"id"`
	DisplayName   string `db:"display_name"`
	LeaguePlayers LeaguePlayers
}

// GetTopLeaguePlayers func
func (l League) GetTopLeaguePlayers() LeaguePlayers {
	sortedLeaguePlayers := l.sortLeaguePlayers()
	return sortedLeaguePlayers
}

func (l League) sortLeaguePlayers() []LeaguePlayer {
	leaguePlayers := make(LeaguePlayers, len(l.LeaguePlayers))
	copy(leaguePlayers, l.LeaguePlayers)

	sort.Sort(sort.Reverse(leaguePlayers))
	return leaguePlayers
}

// GetStatsByPlayerID func
func (l League) GetStatsByPlayerID(playerID string) (LeagueStats, error) {
	var leagueStats LeagueStats

	leaguePlayer := l.getLeaguePlayerByPlayerID(playerID)
	if leaguePlayer == nil {
		return leagueStats, &errors.LeaguePlayerNotFoundError{
			ID: playerID,
		}
	}

	place := l.getPlaceByPlayerID(playerID)
	leagueStats = LeagueStats{
		LeaguePlayer: *leaguePlayer,
		Place:        place,
	}

	return leagueStats, nil
}

func (l League) getLeaguePlayerByPlayerID(playerID string) *LeaguePlayer {
	for i := range l.LeaguePlayers {
		if l.LeaguePlayers[i].PlayerID == playerID {
			return &l.LeaguePlayers[i]
		}
	}
	return nil
}

func (l League) getPlaceByPlayerID(playerID string) int {
	sortedLeaguePlayers := l.sortLeaguePlayers()

	for i := range sortedLeaguePlayers {
		if sortedLeaguePlayers[i].PlayerID == playerID {
			place := i + 1
			return place
		}
	}
	return -1
}

func (l LeaguePlayers) Len() int {
	return len(l)
}

func (l LeaguePlayers) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l LeaguePlayers) Less(i, j int) bool {
	if l[i].Rating != l[j].Rating {
		return l[i].Rating < l[j].Rating
	}

	if l[i].winCount != l[j].winCount {
		return l[i].winCount < l[j].winCount
	}

	iLoseCount := l[i].matchCount - l[i].winCount
	jLoseCount := l[j].matchCount - l[j].winCount
	return iLoseCount > jLoseCount
}
