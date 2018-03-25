package undo

import (
	"github.com/szokodiakos/r8m8/entity"
)

// UndoMatchOutput struct
type UndoMatchOutput struct {
	ReporterPlayer entity.Player
	LeaguePlayers  []entity.LeaguePlayer
	MatchPlayers   []entity.MatchPlayer
}
