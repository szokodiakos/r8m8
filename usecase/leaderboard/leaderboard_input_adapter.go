package leaderboard

// LeaderboardInputAdapter interface
type LeaderboardInputAdapter interface {
	Handle(interface{}) (Input, error)
}
