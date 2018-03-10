package player

import "github.com/szokodiakos/r8m8/transaction"

// SlackRepository interface
type SlackRepository interface {
	GetMultipleByUserIDs(transaction transaction.Transaction, userIDs []string, teamID string) ([]Slack, error)
	Create(transaction transaction.Transaction, slackPlayer Slack) error
}
