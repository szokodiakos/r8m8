package entity

// MatchPlayer struct
type MatchPlayer struct {
	PlayerID     string
	RatingChange int
	HasWon       bool
	leaguePlayer LeaguePlayer
}

// LeaguePlayer func
func (m MatchPlayer) LeaguePlayer() LeaguePlayer {
	return m.leaguePlayer
}

// NewMatchPlayer factory
func NewMatchPlayer(matchPlayer MatchPlayer, leaguePlayer LeaguePlayer) MatchPlayer {
	matchPlayer.leaguePlayer = leaguePlayer
	return matchPlayer
}
