package rating

// Strategy interface
type Strategy interface {
	Calculate(winnerRatings []int, loserRatings []int) Result
}
