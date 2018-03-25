package entity

import (
	"sort"
)

// LeaguePlayers type
type LeaguePlayers []LeaguePlayer

// League struct
type League struct {
	ID            string `db:"id"`
	DisplayName   string `db:"display_name"`
	LeaguePlayers LeaguePlayers
}

// GetTop10LeaguePlayers func
func (l League) GetTop10LeaguePlayers() LeaguePlayers {
	sort.Sort(sort.Reverse(l.LeaguePlayers))

	if len(l.LeaguePlayers) > 10 {
		return l.LeaguePlayers[:10]
	}
	return l.LeaguePlayers
}

func (l LeaguePlayers) Len() int {
	return len(l)
}
func (l LeaguePlayers) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l LeaguePlayers) Less(i, j int) bool {
	return l[i].Rating < l[j].Rating
}
