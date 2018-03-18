package league

import "github.com/szokodiakos/r8m8/league/model"

// GetLeaderboardOutputAdapter interface
type GetLeaderboardOutputAdapter interface {
	Handle(model.GetLeaderboardOutput) (interface{}, error)
}
