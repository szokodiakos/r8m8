package player

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

func TestAddAnyMissingPlayersUnevenPlayers(t *testing.T) {
	playerRepository := &entity.PlayerRepositoryMemory{}
	playerService := NewService(playerRepository)
	tr := transaction.Transaction{}
	unevenPlayers := []entity.Player{
		entity.Player{},
		entity.Player{},
		entity.Player{},
	}

	err := playerService.AddAnyMissingPlayers(tr, unevenPlayers)

	unevenMatchPlayersError := &errors.UnevenMatchPlayersError{}
	assert.EqualError(t, err, unevenMatchPlayersError.Error())
}

func TestAddAnyMissingPlayersDuplicatedPlayers(t *testing.T) {
	playerRepository := &entity.PlayerRepositoryMemory{}
	playerService := NewService(playerRepository)
	tr := transaction.Transaction{}
	duplicatedPlayers := []entity.Player{
		entity.Player{DisplayName: "Dup"},
		entity.Player{DisplayName: "Dup"},
	}

	err := playerService.AddAnyMissingPlayers(tr, duplicatedPlayers)

	duplicatedPlayerExistsError := &errors.DuplicatedPlayerExistsError{}
	assert.EqualError(t, err, duplicatedPlayerExistsError.Error())
}

func TestAddAnyMissingPlayersNoMissingPlayers(t *testing.T) {
	initialRepoPlayers := []entity.Player{
		entity.Player{ID: "One"},
		entity.Player{ID: "Two"},
		entity.Player{ID: "Three"},
	}
	playerRepository := &entity.PlayerRepositoryMemory{Players: initialRepoPlayers}
	playerService := NewService(playerRepository)
	tr := transaction.Transaction{}
	nonMissingPlayers := []entity.Player{
		entity.Player{ID: "One"},
		entity.Player{ID: "Two"},
	}

	err := playerService.AddAnyMissingPlayers(tr, nonMissingPlayers)
	repoPlayers := playerRepository.Players

	assert.Nil(t, err)
	assert.ElementsMatch(t, initialRepoPlayers, repoPlayers)
}

func TestAddAnyMissingPlayersWithMissingPlayers(t *testing.T) {
	initialRepoPlayers := []entity.Player{
		entity.Player{ID: "One"},
		entity.Player{ID: "Two"},
		entity.Player{ID: "Three"},
	}
	playerRepository := &entity.PlayerRepositoryMemory{Players: initialRepoPlayers}
	playerService := NewService(playerRepository)
	tr := transaction.Transaction{}
	missingPlayers := []entity.Player{
		entity.Player{ID: "Four"},
		entity.Player{ID: "Five"},
	}

	err := playerService.AddAnyMissingPlayers(tr, missingPlayers)
	repoPlayers := playerRepository.Players
	expectedRepoPlayers := append(initialRepoPlayers, missingPlayers...)

	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRepoPlayers, repoPlayers)
}
