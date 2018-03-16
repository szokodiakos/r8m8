package league

import (
	"github.com/szokodiakos/r8m8/player"
)

// League struct
type League struct {
	ID          int64           `db:"id"`
	UniqueName  string          `db:"unique_name"`
	DisplayName string          `db:"display_name"`
	Players     []player.Player `db:"players"`
}
