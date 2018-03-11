package player

import "github.com/szokodiakos/r8m8/transaction"

// Service interface
type Service interface {
	GetOrAddPlayers(transaction transaction.Transaction, players []Player) ([]RepoPlayer, error)
	UpdateRatingsForMultiple(transaction transaction.Transaction, repoPlayers []RepoPlayer) error
}

type playerService struct {
	playerRepository Repository
}

func (p *playerService) GetOrAddPlayers(transaction transaction.Transaction, players []Player) ([]RepoPlayer, error) {
	uniqueNames := mapPlayersUniqueNames(players)

	repoPlayers, err := p.playerRepository.GetMultipleByUniqueName(transaction, uniqueNames)
	if err != nil {
		return repoPlayers, err
	}

	if isPlayerMissingFromRepository(players, repoPlayers) {
		missingPlayers := getMissingPlayers(players, repoPlayers)
		err := p.addMultiple(transaction, missingPlayers)
		if err != nil {
			return repoPlayers, err
		}

		repoPlayers, err = p.playerRepository.GetMultipleByUniqueName(transaction, uniqueNames)
		if err != nil {
			return repoPlayers, err
		}
	}
	return repoPlayers, nil
}

func mapPlayersUniqueNames(players []Player) []string {
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

func (p *playerService) addMultiple(transaction transaction.Transaction, players []Player) error {
	for i := range players {
		err := p.playerRepository.Create(transaction, players[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *playerService) UpdateRatingsForMultiple(transaction transaction.Transaction, repoPlayers []RepoPlayer) error {
	for i := range repoPlayers {
		if err := p.playerRepository.UpdateRatingByID(transaction, repoPlayers[i].ID, repoPlayers[i].Rating); err != nil {
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
