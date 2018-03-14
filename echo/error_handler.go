package echo

// ErrorHandler interface
type ErrorHandler interface {
	HandleError(err error) interface{}
}
