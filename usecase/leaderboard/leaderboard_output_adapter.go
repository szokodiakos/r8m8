package leaderboard

// LeaderboardOutputAdapter interface
type LeaderboardOutputAdapter interface {
	Handle(Output, error) (interface{}, error)
}
