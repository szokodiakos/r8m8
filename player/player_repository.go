package player

// Repository interface
type Repository interface {
	Create() (int64, error)
	UpdateRatingByID(ID int64, rating int) error
}
