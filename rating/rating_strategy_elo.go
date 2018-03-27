package rating

import "math"

type strategyElo struct{}

func (s *strategyElo) Calculate(winnerRatings []int, loserRatings []int) Result {
	isWinner := true
	adjustedWinnerRatings := make([]int, len(winnerRatings))
	for i := range winnerRatings {
		adjustedWinnerRatings[i] = calculateRating(winnerRatings[i], loserRatings, isWinner)
	}

	adjustedLoserRatings := make([]int, len(loserRatings))
	for i := range loserRatings {
		adjustedLoserRatings[i] = calculateRating(loserRatings[i], loserRatings, !isWinner)
	}

	return Result{
		WinnerRatings: adjustedWinnerRatings,
		LoserRatings:  adjustedLoserRatings,
	}
}

func calculateRating(rating int, enemyRatings []int, isWinner bool) int {
	ratingChanges := make([]float64, len(enemyRatings))
	for i := range enemyRatings {
		ratingChanges[i] = calculateRatingChange(rating, enemyRatings[i], isWinner)
	}

	averageRatingChange := getAverageInt(ratingChanges)
	return rating + averageRatingChange
}

func calculateRatingChange(rating int, enemyRating int, isWinner bool) float64 {
	const k = 32
	const d = float64(400)
	var s float64
	if isWinner {
		s = 1
	} else {
		s = 0
	}

	ratingDiff := float64(enemyRating - rating)
	e := 1 / (1 + math.Pow(10, (ratingDiff/d)))
	ratingChange := k * (s - e)
	return ratingChange
}

func getAverageInt(numbers []float64) int {
	sum := float64(0)
	for i := range numbers {
		sum += numbers[i]
	}
	average := sum / float64(len(numbers))
	return int(average)
}

// NewStrategyElo factory
func NewStrategyElo() Strategy {
	return &strategyElo{}
}
