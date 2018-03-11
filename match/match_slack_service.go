package match

import (
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

	err = m.matchService.Add(transaction, players)
	if err != nil {
		m.transactionService.Rollback(transaction)
		return messageResponse, err
	}

	err = m.transactionService.Commit(transaction)
	messageResponse = m.slackService.CreateMessageResponse("Success")
	return messageResponse, err
}

// NewSlackService factory
func NewSlackService(matchService Service, slackService slack.Service, playerSlackService player.SlackService, transactionService transaction.Service) SlackService {
	return &matchSlackService{
		matchService:       matchService,
		slackService:       slackService,
		playerSlackService: playerSlackService,
		transactionService: transactionService,
	}
}
