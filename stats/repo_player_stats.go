package stats

// RepoPlayerStats struct
type RepoPlayerStats struct {
	DisplayName string `db:"display_name"`
	Rating      int    `db:"rating"`
	WinCount    int    `db:"won_match_count"`
	MatchCount  int    `db:"total_match_count"`
}
