package match

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/rating"
)

// Service interface
type Service interface {
	CalculatePlayerChanges(repoLeaguePlayers []entity.LeaguePlayer, players []entity.Player) ([]entity.LeaguePlayer, []entity.MatchPlayer)
}

type matchService struct {
	ratingStrategy rating.Strategy
}

func (m *matchService) CalculatePlayerChanges(repoLeaguePlayers []entity.LeaguePlayer, players []entity.Player) ([]entity.LeaguePlayer, []entity.MatchPlayer) {
	winnerPlayers := getWinnerPlayers(players)
	loserPlayers := getLoserPlayers(players)

	winnerRepoLeaguePlayers := getParticipatingRepoLeaguePlayers(winnerPlayers, repoLeaguePlayers)
	loserRepoLeaguePlayers := getParticipatingRepoLeaguePlayers(loserPlayers, repoLeaguePlayers)

	winnerRatings := mapToRatings(winnerRepoLeaguePlayers)
	loserRatings := mapToRatings(loserRepoLeaguePlayers)
	ratingResult := m.ratingStrategy.Calculate(winnerRatings, loserRatings)

	hasWon := true
	hasLost := !hasWon

	adjustedWinnerLeaguePlayers := adjustLeaguePlayers(winnerRepoLeaguePlayers, ratingResult.WinnerRatings)
	adjustedLoserLeaguePlayers := adjustLeaguePlayers(loserRepoLeaguePlayers, ratingResult.LoserRatings)
	adjustedLeaguePlayers := append(adjustedWinnerLeaguePlayers, adjustedLoserLeaguePlayers...)
	mergedLeaguePlayers := mergeInAdjustedLeaguePlayers(adjustedLeaguePlayers, repoLeaguePlayers)

	winnerMatchPlayers := createMatchPlayers(winnerRepoLeaguePlayers, ratingResult.WinnerRatings, hasWon)
	loserMatchPlayers := createMatchPlayers(loserRepoLeaguePlayers, ratingResult.LoserRatings, hasLost)
	matchPlayers := append(winnerMatchPlayers, loserMatchPlayers...)

	return mergedLeaguePlayers, matchPlayers
}

func getWinnerPlayers(players []entity.Player) []entity.Player {
	return players[:(len(players) / 2)]
}

func getLoserPlayers(players []entity.Player) []entity.Player {
	return players[(len(players) / 2):]
}

func getParticipatingRepoLeaguePlayers(players []entity.Player, repoLeaguePlayers []entity.LeaguePlayer) []entity.LeaguePlayer {
	participatingRepoLeaguePlayers := make([]entity.LeaguePlayer, len(players))
	for i := range players {
		for j := range repoLeaguePlayers {
			if players[i].ID == repoLeaguePlayers[j].PlayerID {
				participatingRepoLeaguePlayers[i] = repoLeaguePlayers[j]
			}
		}
	}
	return participatingRepoLeaguePlayers
}

func mapToRatings(leaguePlayers []entity.LeaguePlayer) []int {
	ratings := make([]int, len(leaguePlayers))
	for i := range ratings {
		ratings[i] = leaguePlayers[i].Rating
	}
	return ratings
}

func adjustLeaguePlayers(repoLeaguePlayers []entity.LeaguePlayer, adjustedRatings []int) []entity.LeaguePlayer {
	adjustedRepoLeaguePlayers := make([]entity.LeaguePlayer, len(repoLeaguePlayers))
	for i := range repoLeaguePlayers {
		adjustedRepoLeaguePlayers[i] = repoLeaguePlayers[i]
		adjustedRepoLeaguePlayers[i].Rating = adjustedRatings[i]
	}
	return adjustedRepoLeaguePlayers
}

func createMatchPlayers(repoLeaguePlayers []entity.LeaguePlayer, adjustedRatings []int, hasWon bool) []entity.MatchPlayer {
	matchPlayers := make([]entity.MatchPlayer, len(repoLeaguePlayers))
	for i := range repoLeaguePlayers {
		matchPlayer := entity.MatchPlayer{
			RatingChange: adjustedRatings[i] - repoLeaguePlayers[i].Rating,
			PlayerID:     repoLeaguePlayers[i].PlayerID,
			HasWon:       hasWon,
		}
		matchPlayers[i] = matchPlayer
	}
	return matchPlayers
}

func mergeInAdjustedLeaguePlayers(adjustedLeaguePlayers []entity.LeaguePlayer, repoLeaguePlayers []entity.LeaguePlayer) []entity.LeaguePlayer {
	mergedLeaguePlayers := repoLeaguePlayers
	for i := range adjustedLeaguePlayers {
		for j := range repoLeaguePlayers {
			if adjustedLeaguePlayers[i].PlayerID == repoLeaguePlayers[j].PlayerID {
				mergedLeaguePlayers[j] = adjustedLeaguePlayers[i]
			}
		}
	}
	return mergedLeaguePlayers
}

// NewService creates a service
func NewService(
	ratingStrategy rating.Strategy,
) Service {
	return &matchService{
		ratingStrategy: ratingStrategy,
	}
}
