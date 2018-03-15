package rating

import (
	"github.com/szokodiakos/r8m8/details"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	UpdateRatings(tr transaction.Transaction, repoPlayerIDs []int64, matchID int64) error
}

type ratingService struct {
	strategy          Strategy
	ratingRepository  Repository
	detailsRepository details.Repository
}

func (r *ratingService) UpdateRatings(tr transaction.Transaction, repoPlayerIDs []int64, matchID int64) error {
	repoRatings, err := r.ratingRepository.GetMultipleByPlayerIDs(tr, repoPlayerIDs)
	if err != nil {
		return err
	}
	winnerRepoRatings := getWinnerRepoRatings(repoRatings, repoPlayerIDs)
	loserRepoRatings := getLoserRepoRatings(repoRatings, repoPlayerIDs)

	winnerRatings := mapToRatings(winnerRepoRatings)
	loserRatings := mapToRatings(loserRepoRatings)
	adjustedWinnerRatings, adjustedLoserRatings := r.strategy.Calculate(winnerRatings, loserRatings)

	err = r.adjustRatings(tr, winnerRepoRatings, adjustedWinnerRatings, matchID, true)
	if err != nil {
		return err
	}

	err = r.adjustRatings(tr, loserRepoRatings, adjustedLoserRatings, matchID, false)
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
	loserRepoPlayerIDs := repoPlayerIDs[(len(repoPlayerIDs) / 2):]
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

func (r *ratingService) adjustRatings(tr transaction.Transaction, repoRatings []RepoRating, adjustedRatings []int, matchID int64, hasWon bool) error {
	for i := range repoRatings {
		rating := Rating{
			LeagueID: repoRatings[i].LeagueID,
			PlayerID: repoRatings[i].PlayerID,
			Rating:   adjustedRatings[i],
		}
		err := r.ratingRepository.UpdateRating(tr, rating)
		if err != nil {
			return err
		}

		details := details.Details{
			PlayerID:     repoRatings[i].PlayerID,
			MatchID:      matchID,
			RatingChange: adjustedRatings[i] - repoRatings[i].Rating,
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
