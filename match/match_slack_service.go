package match

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
	"github.com/szokodiakos/r8m8/transaction"
)

// SlackService interface
type SlackService interface {
	Add(values string) (slack.MessageResponse, error)
}

type matchSlackService struct {
	matchService       Service
	slackService       slack.Service
	playerSlackService player.SlackService
	leagueSlackService league.SlackService
	transactionService transaction.Service
}

func (m *matchSlackService) Add(values string) (slack.MessageResponse, error) {
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

	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := m.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := m.playerSlackService.ToPlayer(teamID, userID, userName)
	transaction, err := m.transactionService.Start()
	if err != nil {
		return messageResponse, err
	}

	err = m.matchService.Add(transaction, players, league, reporterPlayer)
	if err != nil {
		m.transactionService.Rollback(transaction)
		return messageResponse, err
	}

	err = m.transactionService.Commit(transaction)

	messageResponse = getSuccessMessageResponse()
	return messageResponse, err
}

func getSuccessMessageResponse() slack.MessageResponse {
	return slack.CreateChannelResponse("Success")
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
