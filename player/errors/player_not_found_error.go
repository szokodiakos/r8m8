package errors

import "fmt"

// PlayerNotFoundError struct
type PlayerNotFoundError struct {
	ID string
}

func (e *PlayerNotFoundError) Error() string {
	return fmt.Sprintf("Player Not Found By ID: %s", e.ID)
}
