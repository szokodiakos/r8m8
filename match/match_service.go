package match

import (
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/player"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	CreateWithPlayers(tr transaction.Transaction, match model.Match, players []playerModel.Player) (model.Match, error)
}

type matchService struct {
	matchRepository       Repository
	matchPlayerRepository PlayerRepository
	playerService         player.Service
	ratingService         rating.Service
}

func (m *matchService) CreateWithPlayers(tr transaction.Transaction, match model.Match, players []playerModel.Player) (model.Match, error) {
	
	// TODO use league players here
	
	// repoLeaguePlayers := leaguePlayerService.GetMultipleByPlayers(tr, players)
	// winnerRepoLeaguePlayers, loserRepoLeaguePlayers...
	// calculate adjusted ratings
	// update league players ratings
	// create match with winner and loser players

	repoPlayers, err := m.playerService.GetRepoPlayersInOrder(tr, players)
	if err != nil {
		return match, err
	}

	winnerRepoPlayers := getWinnerPlayers(repoPlayers)
	loserRepoPlayers := getLoserPlayers(repoPlayers)

	winnerMatchPlayers := 
	matchPlayers, err := m.createMultipleMatchPlayers(tr, repoPlayers)

	repoMatch, err := m.matchRepository.Create(tr, match)
	if err != nil {
		return match, err
	}

	// playerIDs := mapToIDs(repoPlayers)
	// err = m.ratingService.UpdatePlayers(tr, playerIDs, repoMatch.ID)
	// if err != nil {
	// 	return match, err
	// }

	// matchPlayers, err := m.matchPlayerRepository.GetMultipleByMatchID(tr, matchID)
	// if err != nil {
	// 	return match, err
	// }

	winnerMatchPlayers, loserMatchPlayers := sortMatchPlayers(matchPlayers)
	match.WinnerMatchPlayers = winnerMatchPlayers
	match.LoserMatchPlayers = loserMatchPlayers

	return match, nil
}

func getWinnerPlayers(players []playerModel.Player) []playerModel.Player {
	return players[:(len(playerIDs) / 2)]
}

func getLoserPlayers(players []playerModel.Player) []playerModel.Player {
	return players[(len(playerIDs) / 2):]
}

func (m *matchService) createMultipleMatchPlayers(tr transaction.Transaction, players []playerModel.Player) ([]model.MatchPlayer, error) {
	matchPlayers := make([]model.MatchPlayer, len(players))
	for i := range players {
		matchPlayers[i] = model.MatchPlayer{
			Player: players[i],
		}
	}
}

func sortMatchPlayers(matchPlayers []model.MatchPlayer) ([]model.MatchPlayer, []model.MatchPlayer) {
	winnerMatchPlayers := []model.MatchPlayer{}
	loserMatchPlayers := []model.MatchPlayer{}
	for i := range matchPlayers {
		if matchPlayers[i].Details.HasWon == true {
			winnerMatchPlayers = append(winnerMatchPlayers, matchPlayers[i])
		} else {
			loserMatchPlayers = append(loserMatchPlayers, matchPlayers[i])
		}
	}
	return winnerMatchPlayers, loserMatchPlayers
}

// NewService creates a service
func NewService(
	matchRepository Repository,
	matchPlayerRepository PlayerRepository,
	playerService player.Service,
	ratingService rating.Service,
) Service {
	return &matchService{
		matchRepository:       matchRepository,
		matchPlayerRepository: matchPlayerRepository,
		playerService:         playerService,
		ratingService:         ratingService,
	}
}
