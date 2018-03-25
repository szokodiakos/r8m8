package entity

// LeaguePlayer struct
type LeaguePlayer struct {
	PlayerID   string
	Rating     int
	winCount   int
	matchCount int
	player     Player
}

// WinCount func
func (l LeaguePlayer) WinCount() int {
	return l.winCount
}

// MatchCount func
func (l LeaguePlayer) MatchCount() int {
	return l.matchCount
}

// Player func
func (l LeaguePlayer) Player() Player {
	return l.player
}
