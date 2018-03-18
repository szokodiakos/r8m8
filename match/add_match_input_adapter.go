package match

import "github.com/szokodiakos/r8m8/match/model"

// AddMatchInputAdapter interface
type AddMatchInputAdapter interface {
	Handle(interface{}) (model.AddMatchInput, error)
}
