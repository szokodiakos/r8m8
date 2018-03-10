package player

// SlackRepository interface
type SlackRepository interface {
	GetMultipleByUserIDs(userIDs []string, teamID string) ([]Slack, error)
	Create(slackPlayer Slack) error
}
