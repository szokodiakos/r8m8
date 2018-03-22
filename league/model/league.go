package model

// League struct
type League struct {
	ID          int64          `db:"id"`
	UniqueName  string         `db:"unique_name"`
	DisplayName string         `db:"display_name"`
	top10       []LeaguePlayer `db:"top_10_league_players"`
}

// Top10 func
func (l League) Top10() []LeaguePlayer {
	return l.top10
}

// SetTop10 func
func (l *League) SetTop10(top10 []LeaguePlayer) {
	l.top10 = top10
}
