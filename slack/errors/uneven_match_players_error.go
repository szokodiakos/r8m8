package errors

type unevenMatchPlayersError struct {
}

// NewUnevenMatchPlayersError creates an error
func NewUnevenMatchPlayersError() error {
	return &unevenMatchPlayersError{}
}

func (e *unevenMatchPlayersError) Error() string {
	return "Uneven Match Players"
}
