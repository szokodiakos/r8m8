package player

import (
	"github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAddPlayers(tr transaction.Transaction, players []Player, leagueID int64) ([]RepoPlayer, error)
}

type playerService struct {
	playerRepository Repository
	ratingRepository rating.Repository
	initialRating    int
}

func (p *playerService) GetOrAddPlayers(tr transaction.Transaction, players []Player, leagueID int64) ([]RepoPlayer, error) {
	uniqueNames := mapToUniqueNames(players)

	repoPlayers, err := p.playerRepository.GetMultipleByUniqueNames(tr, uniqueNames)
	if err != nil {
		return repoPlayers, err
	}

	if isPlayerMissingFromRepository(players, repoPlayers) {
		missingPlayers := getMissingPlayers(players, repoPlayers)
		err := p.addMultiple(tr, missingPlayers, leagueID)
		if err != nil {
			return repoPlayers, err
		}

		repoPlayers, err = p.playerRepository.GetMultipleByUniqueNames(tr, uniqueNames)
		if err != nil {
			return repoPlayers, err
		}
	}

	repoPlayers, err = sortRepoPlayersByUniqueNames(repoPlayers, uniqueNames)
	return repoPlayers, err
}

func mapToUniqueNames(players []Player) []string {
	uniqueNames := make([]string, len(players))
	for i := range players {
		uniqueNames[i] = players[i].UniqueName
	}
	return uniqueNames
}

func isPlayerMissingFromRepository(players []Player, repoPlayers []RepoPlayer) bool {
	return (len(players) != len(repoPlayers))
}

func getMissingPlayers(players []Player, repoPlayers []RepoPlayer) []Player {
	missingPlayers := make([]Player, 0, len(repoPlayers))

	for i := range players {
		repoPlayer := getRepoCounterpart(players[i], repoPlayers)

		if repoPlayer == (RepoPlayer{}) {
			missingPlayers = append(missingPlayers, players[i])
		}
	}
	return missingPlayers
}

func getRepoCounterpart(player Player, repoPlayers []RepoPlayer) RepoPlayer {
	var repoPlayer RepoPlayer

	for i := range repoPlayers {
		if player.UniqueName == repoPlayers[i].UniqueName {
			repoPlayer = repoPlayers[i]
		}
	}

	return repoPlayer
}

func (p *playerService) addMultiple(tr transaction.Transaction, players []Player, leagueID int64) error {
	for i := range players {
		playerID, err := p.playerRepository.Create(tr, players[i])
		if err != nil {
			return err
		}

		rating := rating.Rating{
			PlayerID: playerID,
			LeagueID: leagueID,
			Rating:   p.initialRating,
		}
		err = p.ratingRepository.Create(tr, rating)
		if err != nil {
			return err
		}
	}

	return nil
}

func sortRepoPlayersByUniqueNames(repoPlayers []RepoPlayer, uniqueNames []string) ([]RepoPlayer, error) {
	orderedRepoPlayers := make([]RepoPlayer, len(repoPlayers))
	for i := range uniqueNames {
		repoPlayer, err := getRepoPlayerByUniqueName(repoPlayers, uniqueNames[i])
		if err != nil {
			return orderedRepoPlayers, err
		}
		orderedRepoPlayers[i] = repoPlayer
	}
	return orderedRepoPlayers, nil
}

func getRepoPlayerByUniqueName(repoPlayers []RepoPlayer, uniqueName string) (RepoPlayer, error) {
	for i := range repoPlayers {
		if repoPlayers[i].UniqueName == uniqueName {
			return repoPlayers[i], nil
		}
	}
	return RepoPlayer{}, &errors.RepoPlayerNotFoundByUniqueNameError{
		UniqueName: uniqueName,
	}
}

// NewService factory
func NewService(playerRepository Repository, ratingRepository rating.Repository, initialRating int) Service {
	return &playerService{
		playerRepository: playerRepository,
		ratingRepository: ratingRepository,
		initialRating:    initialRating,
	}
}
