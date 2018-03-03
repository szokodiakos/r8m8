package match

// Repository interface
type Repository interface {
	Create() (int64, error)
}
