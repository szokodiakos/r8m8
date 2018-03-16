package player

import (
	"github.com/szokodiakos/r8m8/details"
	"github.com/szokodiakos/r8m8/rating"
)

// Player struct
type Player struct {
	ID              int64  `db:"id"`
	UniqueName      string `db:"unique_name"`
	DisplayName     string `db:"display_name"`
	rating.Rating   `db:"rating"`
	details.Details `db:"details"`
}
