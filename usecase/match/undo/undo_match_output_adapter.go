package undo

// OutputAdapter interface
type OutputAdapter interface {
	Handle(Output, error) (interface{}, error)
}
