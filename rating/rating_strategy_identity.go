package rating

type strategyIdentity struct{}

func (s *strategyIdentity) Calculate(winnerRatings []int, loserRatings []int) Result {
	return Result{
		WinnerRatings: winnerRatings,
		LoserRatings:  loserRatings,
	}
}

// NewStrategyIdentity factory
func NewStrategyIdentity() Strategy {
	return &strategyIdentity{}
}
