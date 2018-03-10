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

func (mss *matchSlackService) AddMatch(values string) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse
	requestValues, err := mss.slackService.ParseRequestValues(values)
	if err != nil {
		return messageResponse, err
	}

	text := requestValues.Text
	teamID := requestValues.TeamID

	transaction, err := mss.transactionService.Start()
	if err != nil {
		return messageResponse, err
	}

	slackPlayers, err := mss.playerSlackService.GetOrAddSlackPlayers(transaction, text, teamID)
	if err != nil {
		mss.transactionService.Rollback(transaction)
		return messageResponse, err
	}

	players := mss.getPlayers(slackPlayers)
	err = mss.matchService.Add(transaction, players)
	if err != nil {
		mss.transactionService.Rollback(transaction)
		return messageResponse, err
	}

	err = mss.transactionService.Commit(transaction)
	messageResponse = mss.slackService.CreateMessageResponse("Success")
	return messageResponse, err
}

func (mss *matchSlackService) getPlayers(slackPlayers []player.Slack) []player.Player {
	players := make([]player.Player, len(slackPlayers))

	for i := range slackPlayers {
		players[i] = slackPlayers[i].Player
	}

	return players
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
