package errors

import "fmt"

// PlayerNotFoundByUniqueNameError struct
type PlayerNotFoundByUniqueNameError struct {
	UniqueName string
}

func (e *PlayerNotFoundByUniqueNameError) Error() string {
	return fmt.Sprintf("Repo Player Not Found By Unique Name: %s", e.UniqueName)
}
