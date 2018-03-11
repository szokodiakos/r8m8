package rating

import (
	"github.com/szokodiakos/r8m8/details"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	UpdateRatings(transaction transaction.Transaction, repoPlayerIDs []int64, matchID int64) error
}

type ratingService struct {
	strategy          Strategy
	ratingRepository  Repository
	detailsRepository details.Repository
}

func (r *ratingService) UpdateRatings(transaction transaction.Transaction, repoPlayerIDs []int64, matchID int64) error {
	repoRatings, err := r.ratingRepository.GetMultipleByPlayerIDs(transaction, repoPlayerIDs)
	if err != nil {
		return err
	}
	winnerRepoRatings := getWinnerRepoRatings(repoRatings, repoPlayerIDs)
	loserRepoRatings := getLoserRepoRatings(repoRatings, repoPlayerIDs)

	winnerRatings := mapToRatings(winnerRepoRatings)
	loserRatings := mapToRatings(loserRepoRatings)
	adjustedWinnerRatings, adjustedLoserRatings := r.strategy.Calculate(winnerRatings, loserRatings)

	err = r.adjustRatings(transaction, winnerRepoRatings, adjustedWinnerRatings, matchID)
	if err != nil {
		return err
	}

	err = r.adjustRatings(transaction, loserRepoRatings, adjustedLoserRatings, matchID)
	if err != nil {
		return err
	}

	return nil
}

func getWinnerRepoRatings(repoRatings []RepoRating, repoPlayerIDs []int64) []RepoRating {
	winnerRepoPlayerIDs := repoPlayerIDs[:(len(repoPlayerIDs) / 2)]
	return getRepoRatingsByRepoPlayerIDs(repoRatings, winnerRepoPlayerIDs)
}

func getLoserRepoRatings(repoRatings []RepoRating, repoPlayerIDs []int64) []RepoRating {
	loserRepoPlayerIDs := repoPlayerIDs[:(len(repoPlayerIDs) / 2)]
	return getRepoRatingsByRepoPlayerIDs(repoRatings, loserRepoPlayerIDs)
}

func getRepoRatingsByRepoPlayerIDs(repoRatings []RepoRating, repoPlayerIDs []int64) []RepoRating {
	requestedRepoRatings := make([]RepoRating, 0, len(repoPlayerIDs))
	for _, repoRating := range repoRatings {
		for _, repoPlayerID := range repoPlayerIDs {
			if repoRating.PlayerID == repoPlayerID {
				requestedRepoRatings = append(requestedRepoRatings, repoRating)
			}
		}
	}
	return requestedRepoRatings
}

func mapToRatings(repoRatings []RepoRating) []int {
	ratings := make([]int, len(repoRatings))
	for i := range repoRatings {
		ratings[i] = repoRatings[i].Rating
	}
	return ratings
}

func (r *ratingService) adjustRatings(transaction transaction.Transaction, repoRatings []RepoRating, adjustedRatings []int, matchID int64) error {
	for i := range repoRatings {
		rating := Rating{
			LeagueID: repoRatings[i].LeagueID,
			PlayerID: repoRatings[i].PlayerID,
			Rating:   adjustedRatings[i],
		}
		err := r.ratingRepository.UpdateRating(transaction, rating)
		if err != nil {
			return err
		}

		details := details.Details{
			PlayerID:     repoRatings[i].PlayerID,
			MatchID:      matchID,
			RatingChange: adjustedRatings[i] - repoRatings[i].Rating,
		}
		err = r.detailsRepository.Create(transaction, details)
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
