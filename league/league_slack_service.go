package league

import (
	"fmt"

	"github.com/szokodiakos/r8m8/league/model"
)

// SlackService interface
type SlackService interface {
	ToLeague(teamID string, teamDomain string, channelID string, channelName string) model.League
}

type leagueSlackService struct{}

func (l *leagueSlackService) ToLeague(teamID string, teamDomain string, channelID string, channelName string) model.League {
	uniqueName := fmt.Sprintf("slack_%v_%v", teamID, channelID)
	displayName := fmt.Sprintf("%v %v", teamDomain, channelName)

	return model.League{
		UniqueName:  uniqueName,
		DisplayName: displayName,
	}
}

// NewSlackService factory
func NewSlackService() SlackService {
	return &leagueSlackService{}
}
