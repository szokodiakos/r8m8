package leaderboard

// GetLeaderboardInputAdapter interface
type GetLeaderboardInputAdapter interface {
	Handle(interface{}) (GetLeaderboardInput, error)
}
