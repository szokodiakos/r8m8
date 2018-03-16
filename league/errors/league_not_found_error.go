package errors

import (
	"fmt"
)

// LeagueNotFoundError struct
type LeagueNotFoundError struct {
	UniqueName string
}

func (e *LeagueNotFoundError) Error() string {
	return fmt.Sprintf(`League with Unique Name "%v" Not Found.`, e.UniqueName)
}
