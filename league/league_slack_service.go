package league

import "fmt"

// SlackService interface
type SlackService interface {
	ToLeague(teamID string, teamDomain string, channelID string, channelName string) League
}

type leagueSlackService struct{}

func (l *leagueSlackService) ToLeague(teamID string, teamDomain string, channelID string, channelName string) League {
	uniqueName := fmt.Sprintf("slack_%v_%v", teamID, channelID)
	displayName := fmt.Sprintf("%v %v", teamDomain, channelName)

	return League{
		UniqueName:  uniqueName,
		DisplayName: displayName,
	}
}

// NewSlackService factory
func NewSlackService() SlackService {
	return &leagueSlackService{}
}
