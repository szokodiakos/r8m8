package undo

// UndoMatchInputAdapter interface
type UndoMatchInputAdapter interface {
	Handle(interface{}) (UndoMatchInput, error)
}
