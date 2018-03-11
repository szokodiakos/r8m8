package rating

type strategyIdentity struct{}

func (s *strategyIdentity) Calculate(winnerRatings []int, loserRatings []int) ([]int, []int) {
	return winnerRatings, loserRatings
}

// NewStrategyIdentity factory
func NewStrategyIdentity() Strategy {
	return &strategyIdentity{}
}
