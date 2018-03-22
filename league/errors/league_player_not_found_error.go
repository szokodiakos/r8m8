package errors

import (
	"fmt"
)

// LeaguePlayerNotFoundError struct
type LeaguePlayerNotFoundError struct {
	UniqueName string
}

func (e *LeaguePlayerNotFoundError) Error() string {
	return fmt.Sprintf(`League Player with Unique Name "%v" Not Found.`, e.UniqueName)
}
