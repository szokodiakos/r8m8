package leaderboard

// GetLeaderboardOutputAdapter interface
type GetLeaderboardOutputAdapter interface {
	Handle(GetLeaderboardOutput, error) (interface{}, error)
}
