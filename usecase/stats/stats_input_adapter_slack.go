package stats

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
)

type statsInputAdapterSlack struct {
	slackService       slack.Service
	leagueSlackService league.SlackService
	playerSlackService player.SlackService
}

func (s *statsInputAdapterSlack) Handle(data interface{}) (Input, error) {
	var input Input
	values := data.(string)

	requestValues, err := s.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	teamID := requestValues.TeamID
	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := s.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	userID := requestValues.UserID
	userName := requestValues.UserName
	player := s.playerSlackService.ToPlayer(teamID, channelID, userID, userName)

	input = Input{
		League: league,
		Player: player,
	}
	return input, nil
}

// NewInputAdapterSlack factory
func NewInputAdapterSlack(slackService slack.Service, leagueSlackService league.SlackService, playerSlackService player.SlackService) InputAdapter {
	return &statsInputAdapterSlack{
		slackService:       slackService,
		leagueSlackService: leagueSlackService,
		playerSlackService: playerSlackService,
	}
}
