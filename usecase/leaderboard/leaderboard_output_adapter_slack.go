package leaderboard

import (
	"fmt"
	"strings"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/slack"
)

type leaderboardOutputAdapterSlack struct {
}

func (g *leaderboardOutputAdapterSlack) Handle(output Output, err error) (interface{}, error) {
	if err != nil {
		return getErrorMessageResponse(err)
	}

	league := output.League
	return getSuccessMessageResponse(league), nil
}

func getErrorMessageResponse(err error) (slack.MessageResponse, error) {
	switch err.(type) {
	case *errors.LeagueNotFoundError:
		return getLeagueNotFoundResponse(), nil
	default:
		return slack.MessageResponse{}, err
	}
}

func getLeagueNotFoundResponse() slack.MessageResponse {
	text := `
> Darn! Theres no matches in this League yet! Go play some! :hushed:
`

	return slack.CreateDirectResponse(text)
}

func getSuccessMessageResponse(league entity.League) slack.MessageResponse {
	text := fmt.Sprintf(`
:fire: TOP 10 Leaderboard for *%v* :fire:

%v
	`, league.DisplayName, getPlayersStatsTexts(league.LeaguePlayers))
	return slack.CreateChannelResponse(text)
}

func getPlayersStatsTexts(leaguePlayers []entity.LeaguePlayer) string {
	count := len(leaguePlayers)

	playerText := make([]string, count)
	for i := 0; i < count; i++ {
		icon := getIcon(i + 1)
		displayName := leaguePlayers[i].Player().DisplayName
		rating := leaguePlayers[i].Rating
		winCount := leaguePlayers[i].WinCount()
		matchCount := leaguePlayers[i].MatchCount()
		lossCount := matchCount - winCount
		textTemplate := "> *%v*	%v	*%v*	(%v Win / %v Loss)"
		playerText[i] = fmt.Sprintf(textTemplate, icon, displayName, rating, winCount, lossCount)
	}
	return strings.Join(playerText, "\n")
}

func getIcon(place int) string {
	switch place {
	case 1:
		return ":first_place_medal:"
	case 2:
		return ":second_place_medal:"
	case 3:
		return ":third_place_medal:"
	case 4:
		return ":four:"
	case 5:
		return ":five:"
	case 6:
		return ":six:"
	case 7:
		return ":seven:"
	case 8:
		return ":eight:"
	case 9:
		return ":nine:"
	case 10:
		return ":keycap_ten:"
	default:
		return ""
	}
}

// NewLeaderboardOutputAdapterSlack factory
func NewLeaderboardOutputAdapterSlack() LeaderboardOutputAdapter {
	return &leaderboardOutputAdapterSlack{}
}
