package league

// GetLeaderboardOutputAdapter interface
type GetLeaderboardOutputAdapter interface {
	Handle(GetLeaderboardOutput, error) (interface{}, error)
}
