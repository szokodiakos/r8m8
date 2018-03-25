package errors

import (
	"fmt"
)

// BadSlackPlayerFormatError struct
type BadSlackPlayerFormatError struct {
	SlackPlayer string
}

func (e *BadSlackPlayerFormatError) Error() string {
	return fmt.Sprintf("Bad Slack Player Format: %s", e.SlackPlayer)
}
