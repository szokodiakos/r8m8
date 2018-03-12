package match

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
	"github.com/szokodiakos/r8m8/transaction"
)

// SlackService interface
type SlackService interface {
	AddMatch(values string) (slack.MessageResponse, error)
}

type matchSlackService struct {
	matchService       Service
	slackService       slack.Service
	playerSlackService player.SlackService
	leagueSlackService league.SlackService
	transactionService transaction.Service
}

func (m *matchSlackService) AddMatch(values string) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse

	requestValues, err := m.slackService.ParseRequestValues(values)
	if err != nil {
		return messageResponse, err
	}

	text := requestValues.Text
	teamID := requestValues.TeamID

	players, err := m.playerSlackService.ToPlayers(text, teamID)
	if err != nil {
		return messageResponse, err
	}

	transaction, err := m.transactionService.Start()
	if err != nil {
		return messageResponse, err
	}

	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := m.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := m.playerSlackService.ToPlayer(teamID, userID, userName)
	err = m.matchService.Add(transaction, players, league, reporterPlayer)
	if err != nil {
		m.transactionService.Rollback(transaction)
		return m.handleMatchAddError(err)
	}

	err = m.transactionService.Commit(transaction)
	messageResponse = m.slackService.CreateMessageResponse("Success")
	return messageResponse, err
}

func (m *matchSlackService) handleMatchAddError(err error) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse
	switch err.(type) {
	case *errors.ReporterPlayerNotInLeagueError:
		messageResponse = m.getReporterPlayerNotInLeagueResponse()
		return messageResponse, nil
	case *errors.UnevenMatchPlayersError:
		messageResponse = m.getUnevenMatchPlayersResponse()
		return messageResponse, nil
	default:
		return messageResponse, err
	}
}

func (m *matchSlackService) getReporterPlayerNotInLeagueResponse() slack.MessageResponse {
	return m.slackService.CreateMessageResponse(`
> Darn! You must be the participant of at least one match (including this one). :hushed:
> :exclamation: Please play a match before posting! :exclamation:
	`)
}

func (m *matchSlackService) getUnevenMatchPlayersResponse() slack.MessageResponse {
	return m.slackService.CreateMessageResponse(`
> Darn! Reported players are uneven! :hushed:
> :exclamation: Make sure you report even number of players! :exclamation:
	`)
}

// NewSlackService factory
func NewSlackService(matchService Service, slackService slack.Service, playerSlackService player.SlackService, leagueSlackService league.SlackService, transactionService transaction.Service) SlackService {
	return &matchSlackService{
		matchService:       matchService,
		slackService:       slackService,
		playerSlackService: playerSlackService,
		leagueSlackService: leagueSlackService,
		transactionService: transactionService,
	}
}
