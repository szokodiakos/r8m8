package errors

// ReporterPlayerNotInLeagueError struct
type ReporterPlayerNotInLeagueError struct {
}

func (e *ReporterPlayerNotInLeagueError) Error() string {
	return "Reporter Player Not In League"
}
