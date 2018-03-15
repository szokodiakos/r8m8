package stats

// Leaderboard struct
type Leaderboard struct {
	DisplayName string
	Players     []LeaderboardPlayers
}

// LeaderboardPlayers struct
type LeaderboardPlayers struct {
	DisplayName string
	Rating      int
	Win         int
	Loss        int
}
