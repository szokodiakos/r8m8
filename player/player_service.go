package player

import (
	"github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetRepoPlayers(tr transaction.Transaction, players []model.Player) ([]model.Player, error)
	GetRepoPlayersInOrder(tr transaction.Transaction, players []model.Player) ([]model.Player, error)
}

type playerService struct {
	playerRepository Repository
}

func (p *playerService) GetRepoPlayers(tr transaction.Transaction, players []model.Player) ([]model.Player, error) {
	uniqueNames := mapToUniqueNames(players)
	repoPlayers, err := p.playerRepository.GetMultipleByUniqueNames(tr, uniqueNames)
	return repoPlayers, err
}

func mapToUniqueNames(players []model.Player) []string {
	uniqueNames := make([]string, len(players))
	for i := range players {
		uniqueNames[i] = players[i].UniqueName
	}
	return uniqueNames
}

func sortPlayersByUniqueNames(players []model.Player, uniqueNames []string) ([]model.Player, error) {
	orderedPlayers := make([]model.Player, len(players))
	for i := range uniqueNames {
		player, err := getPlayerByUniqueName(players, uniqueNames[i])
		if err != nil {
			return orderedPlayers, err
		}
		orderedPlayers[i] = player
	}
	return orderedPlayers, nil
}

func getPlayerByUniqueName(players []model.Player, uniqueName string) (model.Player, error) {
	for i := range players {
		if players[i].UniqueName == uniqueName {
			return players[i], nil
		}
	}
	return model.Player{}, &errors.PlayerNotFoundError{
		UniqueName: uniqueName,
	}
}

func (p *playerService) GetRepoPlayersInOrder(tr transaction.Transaction, players []model.Player) ([]model.Player, error) {
	repoPlayers, err := p.GetRepoPlayers(tr, players)
	if err != nil {
		return players, err
	}

	uniqueNames := mapToUniqueNames(players)
	orderedRepoPlayers, err := sortPlayersByUniqueNames(repoPlayers, uniqueNames)
	return orderedRepoPlayers, err
}

// NewService factory
func NewService(playerRepository Repository) Service {
	return &playerService{
		playerRepository: playerRepository,
	}
}
