package player

import (
	"github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	AddAnyMissing(tr transaction.Transaction, players []model.Player) error
	MapToUniqueNames(players []model.Player) []string
}

type playerService struct {
	playerRepository Repository
}

func (p *playerService) AddAnyMissing(tr transaction.Transaction, players []model.Player) error {
	uniqueNames := p.MapToUniqueNames(players)
	repoPlayers, err := p.playerRepository.GetMultipleByUniqueNames(tr, uniqueNames)
	if err != nil {
		return err
	}

	if isMissingPlayerExists(players, repoPlayers) {
		missingPlayers := getMissingPlayers(players, repoPlayers)
		err := p.addMultiple(tr, missingPlayers)
		return err
	}

	return nil
}

func (p *playerService) MapToUniqueNames(players []model.Player) []string {
	uniqueNames := make([]string, len(players))
	for i := range players {
		uniqueNames[i] = players[i].UniqueName
	}
	return uniqueNames
}

func isMissingPlayerExists(players []model.Player, repoPlayers []model.Player) bool {
	return (len(players) != len(repoPlayers))
}

func getMissingPlayers(players []model.Player, repoPlayers []model.Player) []model.Player {
	missingPlayers := make([]model.Player, 0, len(repoPlayers))

	for i := range players {
		repoPlayer := getRepoCounterpart(players[i], repoPlayers)

		if repoPlayer == (model.Player{}) {
			missingPlayers = append(missingPlayers, players[i])
		}
	}
	return missingPlayers
}

func getRepoCounterpart(player model.Player, repoPlayers []model.Player) model.Player {
	var repoPlayer model.Player

	for i := range repoPlayers {
		if player.UniqueName == repoPlayers[i].UniqueName {
			repoPlayer = repoPlayers[i]
		}
	}

	return repoPlayer
}

func (p *playerService) addMultiple(tr transaction.Transaction, players []model.Player) error {
	for i := range players {
		_, err := p.playerRepository.Create(tr, players[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// NewService factory
func NewService(playerRepository Repository) Service {
	return &playerService{
		playerRepository: playerRepository,
	}
}
