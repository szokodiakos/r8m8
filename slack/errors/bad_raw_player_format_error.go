package errors

import "fmt"

type badRawPlayerFormatError struct {
	rawPlayer string
}

// NewBadRawPlayerFormatError creates an error
func NewBadRawPlayerFormatError(rawPlayer string) error {
	return &badRawPlayerFormatError{rawPlayer}
}

func (e *badRawPlayerFormatError) Error() string {
	return fmt.Sprintf("Bad Raw Player Format: %s", e.rawPlayer)
}
