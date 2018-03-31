package leaderboard

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/slack"
)

type getLeaderboardInputAdapterSlack struct {
	slackService       slack.Service
	leagueSlackService league.SlackService
}

func (g *getLeaderboardInputAdapterSlack) Handle(data interface{}) (Input, error) {
	var input Input
	values := data.(string)

	requestValues, err := g.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	teamID := requestValues.TeamID
	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := g.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	input = Input{
		League: league,
	}
	return input, nil
}

// NewGetLeaderboardInputAdapterSlack factory
func NewGetLeaderboardInputAdapterSlack(slackService slack.Service, leagueSlackService league.SlackService) GetLeaderboardInputAdapter {
	return &getLeaderboardInputAdapterSlack{
		slackService:       slackService,
		leagueSlackService: leagueSlackService,
	}
}
