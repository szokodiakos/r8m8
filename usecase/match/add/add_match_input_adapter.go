package add

// InputAdapter interface
type InputAdapter interface {
	Handle(interface{}) (Input, error)
}
