package add

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	playerErrors "github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

// UseCase interface
type UseCase interface {
	Handle(input Input) (Output, error)
}

type addMatchUseCase struct {
	transactionService  transaction.Service
	playerService       player.Service
	leagueService       league.Service
	leaguePlayerService league.PlayerService
	matchService        match.Service
	matchRepository     entity.MatchRepository
	playerRepository    entity.PlayerRepository
	leagueRepository    entity.LeagueRepository
}

func (a *addMatchUseCase) Handle(input Input) (output Output, err error) {
	tr, err := a.transactionService.Start()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = a.transactionService.Commit(tr)
			return
		default:
			err = a.transactionService.Rollback(tr, err)
			return
		}
	}()

	repoLeague, err := a.leagueService.GetOrAddLeague(tr, input.League, input.Players)
	if err != nil {
		return
	}

	err = a.playerService.AddAnyMissingPlayers(tr, input.Players)
	if err != nil {
		return
	}

	repoReporterPlayer, err := a.playerRepository.GetByID(tr, input.ReporterPlayer.ID)
	if err != nil {
		switch err.(type) {
		case *playerErrors.PlayerNotFoundError:
			err = &errors.ReporterPlayerNotInLeagueError{}
			return
		default:
			return
		}
	}

	missingLeaguePlayers := a.leaguePlayerService.CreateAnyMissingLeaguePlayers(repoLeague.LeaguePlayers, input.Players)
	repoLeague.LeaguePlayers = append(repoLeague.LeaguePlayers, missingLeaguePlayers...)
	adjustedLeaguePlayers, matchPlayers := a.matchService.CalculatePlayerChanges(repoLeague.LeaguePlayers, input.Players)

	repoLeague.LeaguePlayers = adjustedLeaguePlayers
	err = a.leagueRepository.Update(tr, repoLeague)
	if err != nil {
		return
	}

	match := entity.Match{
		LeagueID:         repoLeague.ID,
		ReporterPlayerID: repoReporterPlayer.ID,
		MatchPlayers:     matchPlayers,
	}

	repoMatch, err := a.matchRepository.Add(tr, match)
	if err != nil {
		return
	}

	output = Output{
		Match: repoMatch,
	}
	return
}

// NewAddMatchUseCase factory
func NewAddMatchUseCase(
	transactionService transaction.Service,
	playerService player.Service,
	leagueService league.Service,
	leaguePlayerService league.PlayerService,
	matchService match.Service,
	matchRepository entity.MatchRepository,
	playerRepository entity.PlayerRepository,
	leagueRepository entity.LeagueRepository,
) UseCase {
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
