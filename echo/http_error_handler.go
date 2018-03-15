package echo

// HTTPErrorHandler interface
type HTTPErrorHandler interface {
	HandleError(err error) (int, interface{})
}
