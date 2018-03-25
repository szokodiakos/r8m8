package errors

import (
	"fmt"
)

// LeagueNotFoundError struct
type LeagueNotFoundError struct {
	ID string
}

func (e *LeagueNotFoundError) Error() string {
	return fmt.Sprintf(`League with ID "%v" Not Found.`, e.ID)
}
