package stats

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/rating"
)

// PlayerStats struct
type PlayerStats struct {
	player.Player `db:"player"`
	rating.Rating `db:"rating"`
	WinCount      int `db:"won_match_count"`
	MatchCount    int `db:"total_match_count"`
}
