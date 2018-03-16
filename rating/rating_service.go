package rating

import (
	"github.com/szokodiakos/r8m8/details"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	UpdateRatings(tr transaction.Transaction, playerIDs []int64, matchID int64) error
}

type ratingService struct {
	strategy          Strategy
	ratingRepository  Repository
	detailsRepository details.Repository
}

func (r *ratingService) UpdateRatings(tr transaction.Transaction, playerIDs []int64, matchID int64) error {
	ratings, err := r.ratingRepository.GetMultipleByPlayerIDs(tr, playerIDs)
	if err != nil {
		return err
	}
	winnerRatings := getWinnerRatings(ratings, playerIDs)
	loserRatings := getLoserRatings(ratings, playerIDs)

	winnerRatingNumbers := mapToRatingNumbers(winnerRatings)
	loserRatingNumbers := mapToRatingNumbers(loserRatings)
	adjustedWinnerRatingNumbers, adjustedLoserRatingNumbers := r.strategy.Calculate(winnerRatingNumbers, loserRatingNumbers)

	err = r.adjustRatings(tr, winnerRatings, adjustedWinnerRatingNumbers, matchID, true)
	if err != nil {
		return err
	}

	err = r.adjustRatings(tr, loserRatings, adjustedLoserRatingNumbers, matchID, false)
	if err != nil {
		return err
	}

	return nil
}

func getWinnerRatings(ratings []Rating, playerIDs []int64) []Rating {
	winnerPlayerIDs := playerIDs[:(len(playerIDs) / 2)]
	return getRatingsByPlayerIDs(ratings, winnerPlayerIDs)
}

func getLoserRatings(ratings []Rating, playerIDs []int64) []Rating {
	loserPlayerIDs := playerIDs[(len(playerIDs) / 2):]
	return getRatingsByPlayerIDs(ratings, loserPlayerIDs)
}

func getRatingsByPlayerIDs(ratings []Rating, playerIDs []int64) []Rating {
	requestedRatings := make([]Rating, 0, len(playerIDs))
	for _, rating := range ratings {
		for _, playerID := range playerIDs {
			if rating.PlayerID == playerID {
				requestedRatings = append(requestedRatings, rating)
			}
		}
	}
	return requestedRatings
}

func mapToRatingNumbers(ratings []Rating) []int {
	ratingNumbers := make([]int, len(ratings))
	for i := range ratings {
		ratingNumbers[i] = ratings[i].Rating
	}
	return ratingNumbers
}

func (r *ratingService) adjustRatings(tr transaction.Transaction, ratings []Rating, adjustedRatingNumbers []int, matchID int64, hasWon bool) error {
	for i := range ratings {
		rating := Rating{
			LeagueID: ratings[i].LeagueID,
			PlayerID: ratings[i].PlayerID,
			Rating:   adjustedRatingNumbers[i],
		}
		err := r.ratingRepository.UpdateRating(tr, rating)
		if err != nil {
			return err
		}

		details := details.Details{
			PlayerID:     ratings[i].PlayerID,
			MatchID:      matchID,
			RatingChange: adjustedRatingNumbers[i] - ratings[i].Rating,
			HasWon:       hasWon,
		}
		err = r.detailsRepository.Create(tr, details)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewService factory
func NewService(strategy Strategy, ratingRepository Repository, detailsRepository details.Repository) Service {
	return &ratingService{
		strategy:          strategy,
		ratingRepository:  ratingRepository,
		detailsRepository: detailsRepository,
	}
}
