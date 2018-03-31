package leaderboard

// GetLeaderboardOutputAdapter interface
type GetLeaderboardOutputAdapter interface {
	Handle(Output, error) (interface{}, error)
}
