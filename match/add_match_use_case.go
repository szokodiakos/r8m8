package match

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	playerErrors "github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

// AddMatchUseCase interface
type AddMatchUseCase interface {
	Handle(input AddMatchInput) (AddMatchOutput, error)
}

type addMatchUseCase struct {
	transactionService  transaction.Service
	playerService       player.Service
	leagueService       league.Service
	leaguePlayerService league.PlayerService
	matchService        Service
	matchRepository     entity.MatchRepository
	playerRepository    entity.PlayerRepository
	leagueRepository    entity.LeagueRepository
}

func (a *addMatchUseCase) Handle(input AddMatchInput) (AddMatchOutput, error) {
	var output AddMatchOutput

	if isPlayerCountUneven(input.Players) {
		return output, &errors.UnevenMatchPlayersError{}
	}

	tr, err := a.transactionService.Start()
	if err != nil {
		return output, err
	}

	repoLeague, err := a.leagueService.GetOrAddLeague(tr, input.League, input.Players)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	err = a.playerService.AddAnyMissingPlayers(tr, input.Players)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	missingLeaguePlayers, err := a.leaguePlayerService.CreateAnyMissingLeaguePlayers(tr, input.League, input.Players)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	repoLeague = appendMissingLeaguePlayers(repoLeague, missingLeaguePlayers)

	repoReporterPlayer, err := a.playerRepository.GetByID(tr, input.ReporterPlayer.ID)
	if err != nil {
		switch err.(type) {
		case *playerErrors.PlayerNotFoundError:
			return output, a.transactionService.Rollback(tr, &errors.ReporterPlayerNotInLeagueError{})
		default:
			return output, a.transactionService.Rollback(tr, err)
		}
	}

	adjustedLeaguePlayers, matchPlayers := a.matchService.CalculatePlayerChanges(repoLeague.LeaguePlayers, input.Players)

	repoLeague.LeaguePlayers = adjustedLeaguePlayers
	err = a.leagueRepository.Update(tr, repoLeague)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	match := entity.Match{
		LeagueID:         repoLeague.ID,
		ReporterPlayerID: repoReporterPlayer.ID,
		MatchPlayers:     matchPlayers,
	}

	repoMatch, err := a.matchRepository.Add(tr, match)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	output = AddMatchOutput{
		Match: repoMatch,
	}
	err = a.transactionService.Commit(tr)
	return output, err
}

func isPlayerCountUneven(players []entity.Player) bool {
	return (len(players) % 2) != 0
}

func appendMissingLeaguePlayers(league entity.League, missingLeaguePlayers []entity.LeaguePlayer) entity.League {
	league.LeaguePlayers = append(league.LeaguePlayers, missingLeaguePlayers...)
	return league
}

// NewAddMatchUseCase factory
func NewAddMatchUseCase(
	transactionService transaction.Service,
	playerService player.Service,
	leagueService league.Service,
	leaguePlayerService league.PlayerService,
	matchService Service,
	matchRepository entity.MatchRepository,
	playerRepository entity.PlayerRepository,
	leagueRepository entity.LeagueRepository,
) AddMatchUseCase {
	return &addMatchUseCase{
		transactionService:  transactionService,
		playerService:       playerService,
		leagueService:       leagueService,
		leaguePlayerService: leaguePlayerService,
		matchService:        matchService,
		matchRepository:     matchRepository,
		playerRepository:    playerRepository,
		leagueRepository:    leagueRepository,
	}
}
