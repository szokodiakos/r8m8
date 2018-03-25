package match

// AddMatchOutputAdapter interface
type AddMatchOutputAdapter interface {
	Handle(AddMatchOutput, error) (interface{}, error)
}
