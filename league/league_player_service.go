package league

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/player"
)

// PlayerService interface
type PlayerService interface {
	CreateAnyMissingLeaguePlayers(repoLeaguePlayers []entity.LeaguePlayer, players []entity.Player) []entity.LeaguePlayer
	UndoRatingChangesForLeaguePlayers(leaguePlayers []entity.LeaguePlayer, matchPlayers []entity.MatchPlayer) []entity.LeaguePlayer
}

type leaguePlayerService struct {
	playerService    player.Service
	playerRepository player.Repository
	initialRating    int
}

func (l *leaguePlayerService) CreateAnyMissingLeaguePlayers(
	repoLeaguePlayers []entity.LeaguePlayer,
	players []entity.Player,
) []entity.LeaguePlayer {
	missingLeaguePlayers := []entity.LeaguePlayer{}

	for i := range players {
		err := testRepoLeaguePlayerMissing(players[i], repoLeaguePlayers)
		switch err.(type) {
		case *errors.LeaguePlayerNotFoundError:
			missingLeaguePlayers = append(missingLeaguePlayers, entity.LeaguePlayer{
				PlayerID: players[i].ID,
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
func NewPlayerService(playerService player.Service, playerRepository player.Repository, initialRating int) PlayerService {
	return &leaguePlayerService{
		playerService:    playerService,
		playerRepository: playerRepository,
		initialRating:    initialRating,
	}
}
