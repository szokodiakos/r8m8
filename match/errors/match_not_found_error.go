package errors

import (
	"fmt"
)

// MatchNotFoundError struct
type MatchNotFoundError struct {
	ID               int64
	ReporterPlayerID string
}

func (e *MatchNotFoundError) Error() string {
	return fmt.Sprintf(`Match with ID "%v" with ReporterPlayerID "%v" Not Found.`, e.ID, e.ReporterPlayerID)
}
