package league

// GetLeaderboardInputAdapter interface
type GetLeaderboardInputAdapter interface {
	Handle(interface{}) (GetLeaderboardInput, error)
}
