package errors

import "fmt"

// PlayerNotFoundError struct
type PlayerNotFoundError struct {
	UniqueName string
}

func (e *PlayerNotFoundError) Error() string {
	return fmt.Sprintf("Player Not Found By Unique Name: %s", e.UniqueName)
}
