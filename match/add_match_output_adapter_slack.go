package match

import (
	"fmt"
	"strings"

	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/slack"
)

type addMatchOutputAdapterSlack struct {
}

func (a *addMatchOutputAdapterSlack) Handle(output model.AddMatchOutput) (interface{}, error) {
	return getSuccessMessageResponse(output.Match), nil
}

func getSuccessMessageResponse(match model.Match) slack.MessageResponse {
	template := `
*%v* recorded a new Match! Good Game! :muscle:

*Winners* :trophy:
%v
*Losers*
%v
	`
	reporterDisplayName := match.ReporterPlayer.DisplayName
	winnerMatchPlayersText := getMatchPlayersText(match.WinnerMatchPlayers)
	loserMatchPlayersText := getMatchPlayersText(match.LoserMatchPlayers)
	text := fmt.Sprintf(template, reporterDisplayName, winnerMatchPlayersText, loserMatchPlayersText)
	return slack.CreateChannelResponse(text)
}

func getMatchPlayersText(matchPlayers []model.MatchPlayer) string {
	texts := []string{}
	for i := range matchPlayers {
		displayName := matchPlayers[i].Player.DisplayName
		ratingChange := matchPlayers[i].Details.RatingChange
		rating := matchPlayers[i].Rating.Rating
		text := fmt.Sprintf("> *%v* %v and is now at *%v*!", displayName, getRatingChangeText(ratingChange), rating)
		texts = append(texts, text)
	}
	return strings.Join(texts, "\n")
}

func getRatingChangeText(ratingChange int) string {
	if ratingChange < 0 {
		return fmt.Sprintf("has lost *%v* rating", (-1)*ratingChange)
	} else if ratingChange > 0 {
		return fmt.Sprintf("has gained *%v* rating", ratingChange)
	} else {
		return fmt.Sprintf("has no rating change")
	}
}

// NewAddMatchOutputAdapterSlack factory
func NewAddMatchOutputAdapterSlack() AddMatchOutputAdapter {
	return &addMatchOutputAdapterSlack{}
}
