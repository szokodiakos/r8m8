package player

import "github.com/szokodiakos/r8m8/transaction"

// Service interface
type Service interface {
	GetOrAddPlayers(transaction transaction.Transaction, players []Player) ([]DBPlayer, error)
	UpdateRatingsForMultiple(transaction transaction.Transaction, dbPlayers []DBPlayer) error
}

type playerService struct {
	playerRepository Repository
}

func (p *playerService) GetOrAddPlayers(transaction transaction.Transaction, players []Player) ([]DBPlayer, error) {
	uniqueNames := mapPlayersUniqueNames(players)

	dbPlayers, err := p.playerRepository.GetMultipleByUniqueName(transaction, uniqueNames)
	if err != nil {
		return dbPlayers, err
	}

	if isPlayerMissingFromRepository(players, dbPlayers) {
		missingPlayers := getMissingPlayers(players, dbPlayers)
		err := p.addMultiple(transaction, missingPlayers)
		if err != nil {
			return dbPlayers, err
		}

		dbPlayers, err = p.playerRepository.GetMultipleByUniqueName(transaction, uniqueNames)
		if err != nil {
			return dbPlayers, err
		}
	}
	return dbPlayers, nil
}

func mapPlayersUniqueNames(players []Player) []string {
	uniqueNames := make([]string, len(players))
	for i := range players {
		uniqueNames[i] = players[i].UniqueName
	}
	return uniqueNames
}

func isPlayerMissingFromRepository(players []Player, dbPlayers []DBPlayer) bool {
	return (len(players) != len(dbPlayers))
}

func getMissingPlayers(players []Player, dbPlayers []DBPlayer) []Player {
	missingPlayers := make([]Player, 0, len(dbPlayers))

	for i := range players {
		dbPlayer := getDBCounterpart(players[i], dbPlayers)

		if dbPlayer == (DBPlayer{}) {
			missingPlayers = append(missingPlayers, players[i])
		}
	}
	return missingPlayers
}

func getDBCounterpart(player Player, dbPlayers []DBPlayer) DBPlayer {
	var dbPlayer DBPlayer

	for i := range dbPlayers {
		if player.UniqueName == dbPlayers[i].UniqueName {
			dbPlayer = dbPlayers[i]
		}
	}

	return dbPlayer
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

func (p *playerService) UpdateRatingsForMultiple(transaction transaction.Transaction, dbPlayers []DBPlayer) error {
	for i := range dbPlayers {
		if err := p.playerRepository.UpdateRatingByID(transaction, dbPlayers[i].ID, dbPlayers[i].Rating); err != nil {
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
