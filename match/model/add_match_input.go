package model

import (
	leagueModel "github.com/szokodiakos/r8m8/league/model"
	playerModel "github.com/szokodiakos/r8m8/player/model"
)

// AddMatchInput struct
type AddMatchInput struct {
	Players        []playerModel.Player
	League         leagueModel.League
	ReporterPlayer playerModel.Player
}
