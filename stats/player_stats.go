package stats

// PlayerStats struct
type PlayerStats struct {
	DisplayName string `db:"display_name"`
	Rating      int    `db:"rating"`
	WinCount    int    `db:"won_match_count"`
	MatchCount  int    `db:"total_match_count"`
}
