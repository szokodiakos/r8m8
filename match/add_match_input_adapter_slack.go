package match

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

func (a *addMatchInputAdapterSlack) Handle(data interface{}) (AddMatchInput, error) {
	values := data.(string)
	var input AddMatchInput

	requestValues, err := a.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	text := requestValues.Text
	teamID := requestValues.TeamID
	players, err := a.playerSlackService.ToPlayers(text, teamID)
	if err != nil {
		return input, err
	}

	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := a.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := a.playerSlackService.ToPlayer(teamID, userID, userName)

	input = AddMatchInput{
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
) AddMatchInputAdapter {
	return &addMatchInputAdapterSlack{
		slackService:       slackService,
		playerSlackService: playerSlackService,
		leagueSlackService: leagueSlackService,
	}
}
