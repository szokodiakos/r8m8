package league

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerService interface
type PlayerService interface {
	CreateAnyMissingLeaguePlayers(tr transaction.Transaction, league entity.League, players []entity.Player) ([]entity.LeaguePlayer, error)
	UndoRatingChangesForLeaguePlayers(leaguePlayers []entity.LeaguePlayer, matchPlayers []entity.MatchPlayer) []entity.LeaguePlayer
}

type leaguePlayerService struct {
	playerService    player.Service
	playerRepository entity.PlayerRepository
	initialRating    int
}

func (l *leaguePlayerService) CreateAnyMissingLeaguePlayers(tr transaction.Transaction, repoLeague entity.League, players []entity.Player) ([]entity.LeaguePlayer, error) {
	ids := l.playerService.MapToIDs(players)
	repoLeaguePlayers := repoLeague.LeaguePlayers
	missingLeaguePlayers := []entity.LeaguePlayer{}

	if isMissingLeaguePlayerExists(players, repoLeaguePlayers) {
		repoPlayers, err := l.playerRepository.GetMultipleByIDs(tr, ids)
		if err != nil {
			return nil, err
		}

		missingLeaguePlayers = l.createMissingLeaguePlayers(repoPlayers, repoLeaguePlayers, repoLeague)
	}

	return missingLeaguePlayers, nil
}

func isMissingLeaguePlayerExists(players []entity.Player, repoLeaguePlayers []entity.LeaguePlayer) bool {
	return (len(players) != len(repoLeaguePlayers))
}

func (l *leaguePlayerService) createMissingLeaguePlayers(repoPlayers []entity.Player, repoLeaguePlayers []entity.LeaguePlayer, league entity.League) []entity.LeaguePlayer {
	missingLeaguePlayers := []entity.LeaguePlayer{}

	for i := range repoPlayers {
		err := testRepoLeaguePlayerMissing(repoPlayers[i], repoLeaguePlayers)
		switch err.(type) {
		case *errors.LeaguePlayerNotFoundError:
			missingLeaguePlayers = append(missingLeaguePlayers, entity.LeaguePlayer{
				PlayerID: repoPlayers[i].ID,
				Rating:   l.initialRating,
			})
		}
	}

	return missingLeaguePlayers
}

func testRepoLeaguePlayerMissing(repoPlayer entity.Player, repoLeaguePlayers []entity.LeaguePlayer) error {
	for i := range repoLeaguePlayers {
		if repoLeaguePlayers[i].PlayerID == repoPlayer.ID {
			return nil
		}
	}

	return &errors.LeaguePlayerNotFoundError{
		ID: repoPlayer.ID,
	}
}

func (l *leaguePlayerService) UndoRatingChangesForLeaguePlayers(leaguePlayers []entity.LeaguePlayer, matchPlayers []entity.MatchPlayer) []entity.LeaguePlayer {
	adjustedLeaguePlayers := make([]entity.LeaguePlayer, len(leaguePlayers))
	copy(adjustedLeaguePlayers, leaguePlayers)

	for i := range matchPlayers {
		for j := range leaguePlayers {
			if isLeaguePlayerParticipatedInMatch(matchPlayers[i], leaguePlayers[j]) {
				adjustedLeaguePlayers[j].Rating -= matchPlayers[i].RatingChange
			}
		}
	}

	return adjustedLeaguePlayers
}

func isLeaguePlayerParticipatedInMatch(matchPlayer entity.MatchPlayer, leaguePlayer entity.LeaguePlayer) bool {
	return matchPlayer.PlayerID == leaguePlayer.PlayerID
}

// NewPlayerService factory
func NewPlayerService(playerService player.Service, playerRepository entity.PlayerRepository, initialRating int) PlayerService {
	return &leaguePlayerService{
		playerService:    playerService,
		playerRepository: playerRepository,
		initialRating:    initialRating,
	}
}
