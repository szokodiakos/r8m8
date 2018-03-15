package stats

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/slack"
)

// SlackService interface
type SlackService interface {
	GetLeaderboard(values string) (slack.MessageResponse, error)
}

type statsSlackService struct {
	statsService       Service
	leagueSlackService league.SlackService
	slackService       slack.Service
}

func (s *statsSlackService) GetLeaderboard(values string) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse

	requestValues, err := s.slackService.ParseRequestValues(values)
	if err != nil {
		return messageResponse, err
	}

	teamID := requestValues.TeamID
	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := s.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	leaderboard, err := s.statsService.GetLeaderboard(league)
	if err != nil {
		return messageResponse, err
	}

	messageResponse = s.getSuccessMessageResponse(leaderboard)
	return messageResponse, nil
}

func (s *statsSlackService) getSuccessMessageResponse(leaderboard Leaderboard) slack.MessageResponse {
	var messageResponse slack.MessageResponse
	return messageResponse
}

// NewSlackService factory
func NewSlackService(statsService Service, leagueSlackService league.SlackService, slackService slack.Service) SlackService {
	return &statsSlackService{
		statsService:       statsService,
		leagueSlackService: leagueSlackService,
		slackService:       slackService,
	}
}
