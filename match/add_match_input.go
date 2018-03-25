package match

import "github.com/szokodiakos/r8m8/entity"

// AddMatchInput struct
type AddMatchInput struct {
	Players        []entity.Player
	League         entity.League
	ReporterPlayer entity.Player
}
