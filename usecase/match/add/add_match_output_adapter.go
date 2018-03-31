package add

// OutputAdapter interface
type OutputAdapter interface {
	Handle(Output, error) (interface{}, error)
}
