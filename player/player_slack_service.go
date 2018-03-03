package player

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// SlackService interface
type SlackService interface {
	GetOrAddSlackPlayers(text string, teamID string) ([]Slack, error)
}

type playerSlackService struct {
	playerSlackRepository    SlackRepository
	transactionService       transaction.Service
	playerService            Service
	playerSlackParserService SlackParserService
}

func (pss *playerSlackService) GetOrAddSlackPlayers(text string, teamID string) ([]Slack, error) {
	slackPlayers, err := pss.playerSlackParserService.Parse(text, teamID)
	if err != nil {
		return slackPlayers, err
	}

	slackPlayerUserIDs := pss.getSlackPlayersUserIDs(slackPlayers)
	repositorySlackPlayers, err := pss.playerSlackRepository.GetMultipleByUserIDs(slackPlayerUserIDs, teamID)
	if err != nil {
		return repositorySlackPlayers, err
	}

	if pss.isSlackPlayerMissingFromRepository(repositorySlackPlayers, slackPlayers) {
		pss.addMissingSlackPlayers(repositorySlackPlayers, slackPlayers)

		repositorySlackPlayers, err = pss.playerSlackRepository.GetMultipleByUserIDs(slackPlayerUserIDs, teamID)
		if err != nil {
			return repositorySlackPlayers, err
		}
	}

	return repositorySlackPlayers, nil
}

func (pss *playerSlackService) getSlackPlayersUserIDs(slackPlayers []Slack) []string {
	slackPlayerUserIDs := make([]string, len(slackPlayers))

	for i := range slackPlayers {
		slackPlayerUserIDs[i] = slackPlayers[i].UserID
	}

	return slackPlayerUserIDs
}

func (pss *playerSlackService) isSlackPlayerMissingFromRepository(repositorySlackPlayers []Slack, slackPlayers []Slack) bool {
	return (len(repositorySlackPlayers) != len(slackPlayers))
}

func (pss *playerSlackService) getMissingSlackPlayers(repositorySlackPlayers []Slack, slackPlayers []Slack) []Slack {
	missingSlackPlayers := make([]Slack, 0, len(slackPlayers))

	for _, slackPlayer := range slackPlayers {
		repositorySlackPlayer := pss.getCounterpart(slackPlayer, repositorySlackPlayers)

		if repositorySlackPlayer == (Slack{}) {
			missingSlackPlayers = append(missingSlackPlayers, repositorySlackPlayer)
		}
	}
	return missingSlackPlayers
}

func (pss *playerSlackService) getCounterpart(slackPlayer Slack, repositorySlackPlayers []Slack) Slack {
	var foundRepositorySlackPlayer Slack

	for _, repositorySlackPlayer := range repositorySlackPlayers {
		if slackPlayer.UserID == repositorySlackPlayer.UserID && slackPlayer.TeamID == repositorySlackPlayer.TeamID {
			foundRepositorySlackPlayer = repositorySlackPlayer
		}
	}

	return foundRepositorySlackPlayer
}

func (pss *playerSlackService) addMultiple(playerIDs []int64, slackPlayers []Slack) error {
	for i, slackPlayer := range slackPlayers {
		slackPlayer.Player = Player{
			ID: playerIDs[i],
		}
		_, err := pss.playerSlackRepository.Create(slackPlayer)

		if err != nil {
			return err
		}
	}

	return nil
}

func (pss *playerSlackService) addMissingSlackPlayers(repositorySlackPlayers []Slack, slackPlayers []Slack) error {
	missingSlackPlayers := pss.getMissingSlackPlayers(repositorySlackPlayers, slackPlayers)
	tr, err := pss.transactionService.Start()
	if err != nil {
		return err
	}

	missingSlackPlayerCount := len(missingSlackPlayers)
	addedPlayerIDs, err := pss.playerService.AddMultiple(missingSlackPlayerCount)
	if err != nil {
		pss.transactionService.Rollback(tr)
		return err
	}

	err = pss.addMultiple(addedPlayerIDs, missingSlackPlayers)
	if err != nil {
		pss.transactionService.Rollback(tr)
		return err
	}

	pss.transactionService.Commit(tr)
	return nil
}

// NewSlackService factory
func NewSlackService(playerSlackRepository SlackRepository, transactionService transaction.Service, playerService Service, playerSlackParserService SlackParserService) SlackService {
	return &playerSlackService{
		playerSlackRepository:    playerSlackRepository,
		transactionService:       transactionService,
		playerService:            playerService,
		playerSlackParserService: playerSlackParserService,
	}
}
