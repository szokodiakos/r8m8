package errors

import "fmt"

type badSlackPlayerFormatError struct {
	slackPlayer string
}

func (e *badSlackPlayerFormatError) Error() string {
	return fmt.Sprintf("Bad Slack Player Format: %s", e.slackPlayer)
}

// NewBadSlackPlayerFormatError factory
func NewBadSlackPlayerFormatError(slackPlayer string) error {
	return &badSlackPlayerFormatError{slackPlayer}
}
