package errors

type unevenMatchPlayersError struct {
}

func (e *unevenMatchPlayersError) Error() string {
	return "Uneven Match Players"
}

// NewUnevenMatchPlayersError factory
func NewUnevenMatchPlayersError() error {
	return &unevenMatchPlayersError{}
}
