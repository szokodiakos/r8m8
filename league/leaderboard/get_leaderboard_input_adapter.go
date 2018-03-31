package leaderboard

// GetLeaderboardInputAdapter interface
type GetLeaderboardInputAdapter interface {
	Handle(interface{}) (Input, error)
}
