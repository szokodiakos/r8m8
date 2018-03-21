package model

// League struct
type League struct {
	ID          int64  `db:"id"`
	UniqueName  string `db:"unique_name"`
	DisplayName string `db:"display_name"`
	Top10       []LeaguePlayer
}
