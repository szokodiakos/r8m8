package undo

// UndoMatchOutputAdapter interface
type UndoMatchOutputAdapter interface {
	Handle(UndoMatchOutput, error) (interface{}, error)
}
