package errors

// InvalidVerificationTokenError struct
type InvalidVerificationTokenError struct {
}

func (e *InvalidVerificationTokenError) Error() string {
	return "Invalid Verification Token Error"
}
