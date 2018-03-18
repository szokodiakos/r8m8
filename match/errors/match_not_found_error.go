package errors

import (
	"fmt"
)

// MatchNotFoundError struct
type MatchNotFoundError struct {
	ID int64
}

func (e *MatchNotFoundError) Error() string {
	return fmt.Sprintf(`Match with ID "%v" Not Found.`, e.ID)
}
