package player

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	AddAnyMissingPlayers(tr transaction.Transaction, players []entity.Player) error
}

type playerService struct {
	playerRepository Repository
}

func (p *playerService) AddAnyMissingPlayers(tr transaction.Transaction, players []entity.Player) error {
	if isPlayerCountUneven(players) {
		return &errors.UnevenMatchPlayersError{}
	}

	if isDuplicatedPlayerExists(players) {
		return &errors.DuplicatedPlayerExistsError{}
	}

	ids := mapToIDs(players)
	repoPlayers, err := p.playerRepository.GetMultipleByIDs(tr, ids)
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

func isPlayerCountUneven(players []entity.Player) bool {
	return (len(players) % 2) != 0
}

func isDuplicatedPlayerExists(players []entity.Player) bool {
	ids := mapToIDs(players)
	idMap := map[string]bool{}

	for _, id := range ids {
		if idMap[id] == false {
			idMap[id] = true
		} else {
			return true
		}
	}

	return false
}

func mapToIDs(players []entity.Player) []string {
	ids := make([]string, len(players))
	for i := range players {
		ids[i] = players[i].ID
	}
	return ids
}

func isMissingPlayerExists(players []entity.Player, repoPlayers []entity.Player) bool {
	return (len(players) != len(repoPlayers))
}

func getMissingPlayers(players []entity.Player, repoPlayers []entity.Player) []entity.Player {
	missingPlayers := []entity.Player{}

	for i := range players {
		repoPlayer := getRepoCounterpart(players[i], repoPlayers)

		if repoPlayer == (entity.Player{}) {
			missingPlayers = append(missingPlayers, players[i])
		}
	}
	return missingPlayers
}

func getRepoCounterpart(player entity.Player, repoPlayers []entity.Player) entity.Player {
	var repoPlayer entity.Player

	for i := range repoPlayers {
		if player.ID == repoPlayers[i].ID {
			repoPlayer = repoPlayers[i]
		}
	}

	return repoPlayer
}

func (p *playerService) addMultiple(tr transaction.Transaction, players []entity.Player) error {
	for i := range players {
		_, err := p.playerRepository.Add(tr, players[i])
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
