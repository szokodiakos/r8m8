package match

import (
	"github.com/szokodiakos/r8m8/league"
	leagueModel "github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	CreateWithPlayerUniqueNames(tr transaction.Transaction, match model.Match, uniqueNames []string) (model.Match, error)
}

type matchService struct {
	ratingStrategy        rating.Strategy
	matchRepository       Repository
	matchPlayerRepository PlayerRepository
	leaguePlayerService   league.PlayerService
}

func (m *matchService) CreateWithPlayerUniqueNames(tr transaction.Transaction, match model.Match, uniqueNames []string) (model.Match, error) {
	repoLeaguePlayers, err := m.leaguePlayerService.GetMultipleByUniqueNamesInOrder(tr, uniqueNames)
	if err != nil {
		return match, err
	}

	winnerRepoLeaguePlayers := getWinnerLeaguePlayers(repoLeaguePlayers)
	loserRepoLeaguePlayers := getLoserLeaguePlayers(repoLeaguePlayers)

	winnerRatings := mapToRatings(winnerRepoLeaguePlayers)
	loserRatings := mapToRatings(loserRepoLeaguePlayers)
	adjustedWinnerRatings, adjustedLoserRatings := m.ratingStrategy.Calculate(winnerRatings, loserRatings)

	adjustedWinnerLeaguePlayers := adjustLeaguePlayersRatings(winnerRepoLeaguePlayers, adjustedWinnerRatings)
	adjustedLoserLeaguePlayers := adjustLeaguePlayersRatings(winnerRepoLeaguePlayers, adjustedLoserRatings)
	adjustedLeaguePlayers := append(adjustedWinnerLeaguePlayers, adjustedLoserLeaguePlayers...)
	err = m.leaguePlayerService.UpdateMultiple(tr, adjustedLeaguePlayers)
	if err != nil {
		return match, err
	}

	hasWon := true
	winnerMatchPlayers := m.createMultipleMatchPlayers(winnerRepoLeaguePlayers, adjustedWinnerLeaguePlayers, hasWon)
	loserMatchPlayers := m.createMultipleMatchPlayers(loserRepoLeaguePlayers, adjustedLoserLeaguePlayers, !hasWon)

	match.MatchPlayers = append(winnerMatchPlayers, loserMatchPlayers...)
	repoMatch, err := m.matchRepository.Create(tr, match)
	if err != nil {
		return match, err
	}

	return repoMatch, nil
}

func getWinnerLeaguePlayers(leaguePlayers []leagueModel.LeaguePlayer) []leagueModel.LeaguePlayer {
	return leaguePlayers[:(len(leaguePlayers) / 2)]
}

func getLoserLeaguePlayers(leaguePlayers []leagueModel.LeaguePlayer) []leagueModel.LeaguePlayer {
	return leaguePlayers[(len(leaguePlayers) / 2):]
}

func mapToRatings(leaguePlayers []leagueModel.LeaguePlayer) []int {
	ratings := make([]int, len(leaguePlayers))
	for i := range ratings {
		ratings[i] = leaguePlayers[i].Rating
	}
	return ratings
}

func adjustLeaguePlayersRatings(repoLeaguePlayers []leagueModel.LeaguePlayer, adjustedRatings []int) []leagueModel.LeaguePlayer {
	adjustedRepoLeaguePlayers := make([]leagueModel.LeaguePlayer, len(repoLeaguePlayers))

	for i := range repoLeaguePlayers {
		adjustedRepoLeaguePlayers[i].Rating = adjustedRatings[i]
	}

	return adjustedRepoLeaguePlayers
}

func (m *matchService) createMultipleMatchPlayers(repoLeaguePlayers []leagueModel.LeaguePlayer, adjustedLeaguePlayers []leagueModel.LeaguePlayer, hasWon bool) []model.MatchPlayer {
	matchPlayers := make([]model.MatchPlayer, len(repoLeaguePlayers))

	for i := range repoLeaguePlayers {
		matchPlayers[i] = model.MatchPlayer{
			Player:       repoLeaguePlayers[i].Player,
			HasWon:       hasWon,
			RatingChange: adjustedLeaguePlayers[i].Rating - repoLeaguePlayers[i].Rating,
		}
	}

	return matchPlayers
}

// NewService creates a service
func NewService(
	ratingStrategy rating.Strategy,
	matchRepository Repository,
	matchPlayerRepository PlayerRepository,
	leaguePlayerService league.PlayerService,
) Service {
	return &matchService{
		ratingStrategy:        ratingStrategy,
		matchRepository:       matchRepository,
		matchPlayerRepository: matchPlayerRepository,
		leaguePlayerService:   leaguePlayerService,
	}
}
