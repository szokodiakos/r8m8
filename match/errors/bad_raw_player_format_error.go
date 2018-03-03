package errors

import "fmt"

type badRawPlayerFormatError struct {
	rawPlayer string
}

func (e *badRawPlayerFormatError) Error() string {
	return fmt.Sprintf("Bad Raw Player Format: %s", e.rawPlayer)
}

// NewBadRawPlayerFormatError factory
func NewBadRawPlayerFormatError(rawPlayer string) error {
	return &badRawPlayerFormatError{rawPlayer}
}
