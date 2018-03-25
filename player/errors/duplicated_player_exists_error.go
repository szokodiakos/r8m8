package errors

// DuplicatedPlayerExistsError struct
type DuplicatedPlayerExistsError struct {
}

func (e *DuplicatedPlayerExistsError) Error() string {
	return "Duplicated Player Exists"
}
