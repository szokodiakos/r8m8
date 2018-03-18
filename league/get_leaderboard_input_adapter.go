package league

import "github.com/szokodiakos/r8m8/league/model"

// GetLeaderboardInputAdapter interface
type GetLeaderboardInputAdapter interface {
	Handle(interface{}) (model.GetLeaderboardInput, error)
}
