package errors

// UnevenMatchPlayersError struct
type UnevenMatchPlayersError struct {
}

func (e *UnevenMatchPlayersError) Error() string {
	return "Uneven Match Players"
}
