package rating

// Rating struct
type Rating struct {
	PlayerID int64 `db:"player_id"`
	LeagueID int64 `db:"league_id"`
	Rating   int   `db:"rating"`
}
