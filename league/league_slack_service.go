package league

import (
	"fmt"

	"github.com/szokodiakos/r8m8/entity"
)

// SlackService interface
type SlackService interface {
	ToLeague(teamID string, teamDomain string, channelID string, channelName string) entity.League
}

type leagueSlackService struct{}

func (l *leagueSlackService) ToLeague(teamID string, teamDomain string, channelID string, channelName string) entity.League {
	id := fmt.Sprintf("slack_%v_%v", teamID, channelID)
	displayName := fmt.Sprintf("%v %v", teamDomain, channelName)

	return entity.League{
		ID:          id,
		DisplayName: displayName,
	}
}

// NewSlackService factory
func NewSlackService() SlackService {
	return &leagueSlackService{}
}
