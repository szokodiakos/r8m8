package match

// DetailsRepository interface
type DetailsRepository interface {
	Create(matchDetails Details) error
}
