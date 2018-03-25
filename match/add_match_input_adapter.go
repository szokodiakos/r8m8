package match

// AddMatchInputAdapter interface
type AddMatchInputAdapter interface {
	Handle(interface{}) (AddMatchInput, error)
}
