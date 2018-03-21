package match

import "github.com/szokodiakos/r8m8/match/model"

// AddMatchOutputAdapter interface
type AddMatchOutputAdapter interface {
	Handle(model.AddMatchOutput, error) (interface{}, error)
}
