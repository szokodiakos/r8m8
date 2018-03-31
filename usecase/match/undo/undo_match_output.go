package undo

import (
	"github.com/szokodiakos/r8m8/entity"
)

// Output struct
type Output struct {
	ReporterPlayer entity.Player
	LeaguePlayers  []entity.LeaguePlayer
	MatchPlayers   []entity.MatchPlayer
}
