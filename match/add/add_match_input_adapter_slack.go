package add

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
)

type addMatchInputAdapterSlack struct {
	slackService       slack.Service
	playerSlackService player.SlackService
	leagueSlackService league.SlackService
}

func (a *addMatchInputAdapterSlack) Handle(data interface{}) (Input, error) {
	values := data.(string)
	var input Input

	requestValues, err := a.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	text := requestValues.Text
	teamID := requestValues.TeamID
	channelID := requestValues.ChannelID
	players, err := a.playerSlackService.ToPlayers(text, teamID, channelID)
	if err != nil {
		return input, err
	}

	teamDomain := requestValues.TeamDomain
	channelName := requestValues.ChannelName
	league := a.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := a.playerSlackService.ToPlayer(teamID, channelID, userID, userName)

	input = Input{
		League:         league,
		Players:        players,
		ReporterPlayer: reporterPlayer,
	}
	return input, err
}

// NewAddMatchInputAdapterSlack factory
func NewAddMatchInputAdapterSlack(
	slackService slack.Service,
	playerSlackService player.SlackService,
	leagueSlackService league.SlackService,
) InputAdapter {
	return &addMatchInputAdapterSlack{
		slackService:       slackService,
		playerSlackService: playerSlackService,
		leagueSlackService: leagueSlackService,
	}
}
