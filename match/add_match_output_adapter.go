package match

import "github.com/szokodiakos/r8m8/match/model"

type AddMatchOutputAdapter interface {
	Handle(model.AddMatchOutput) (interface{}, error)
}
