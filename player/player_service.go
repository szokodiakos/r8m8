package player

import (
	"github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAddPlayers(tr transaction.Transaction, players []Player, leagueID int64) ([]Player, error)
}

type playerService struct {
	playerRepository Repository
	ratingRepository rating.Repository
	initialRating    int
}

func (p *playerService) GetOrAddPlayers(tr transaction.Transaction, players []Player, leagueID int64) ([]Player, error) {
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

	repoPlayers, err = sortPlayersByUniqueNames(repoPlayers, uniqueNames)
	return repoPlayers, err
}

func mapToUniqueNames(players []Player) []string {
	uniqueNames := make([]string, len(players))
	for i := range players {
		uniqueNames[i] = players[i].UniqueName
	}
	return uniqueNames
}

func isPlayerMissingFromRepository(players []Player, repoPlayers []Player) bool {
	return (len(players) != len(repoPlayers))
}

func getMissingPlayers(players []Player, repoPlayers []Player) []Player {
	missingPlayers := make([]Player, 0, len(repoPlayers))

	for i := range players {
		repoPlayer := getRepoCounterpart(players[i], repoPlayers)

		if repoPlayer == (Player{}) {
			missingPlayers = append(missingPlayers, players[i])
		}
	}
	return missingPlayers
}

func getRepoCounterpart(player Player, repoPlayers []Player) Player {
	var repoPlayer Player

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

func sortPlayersByUniqueNames(players []Player, uniqueNames []string) ([]Player, error) {
	orderedPlayers := make([]Player, len(players))
	for i := range uniqueNames {
		player, err := getPlayerByUniqueName(players, uniqueNames[i])
		if err != nil {
			return orderedPlayers, err
		}
		orderedPlayers[i] = player
	}
	return orderedPlayers, nil
}

func getPlayerByUniqueName(players []Player, uniqueName string) (Player, error) {
	for i := range players {
		if players[i].UniqueName == uniqueName {
			return players[i], nil
		}
	}
	return Player{}, &errors.PlayerNotFoundByUniqueNameError{
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
