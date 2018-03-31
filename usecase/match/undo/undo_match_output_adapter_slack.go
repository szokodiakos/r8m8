package undo

import (
	"fmt"
	"strings"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/slack"
)

type undoMatchOutputAdapterSlack struct {
}

func (a *undoMatchOutputAdapterSlack) Handle(output Output, err error) (interface{}, error) {
	if err != nil {
		return getErrorMessageResponse(err)
	}

	return getSuccessMessageResponse(output.ReporterPlayer, output.LeaguePlayers, output.MatchPlayers), nil
}

func getErrorMessageResponse(err error) (slack.MessageResponse, error) {
	switch err.(type) {
	case *errors.MatchNotFoundError:
		return getMatchNotFoundResponse(), nil
	default:
		return slack.MessageResponse{}, err
	}
}

func getMatchNotFoundResponse() slack.MessageResponse {
	text := `
> Darn! You don't have a match yet! :hushed:
`

	return slack.CreateDirectResponse(text)
}

func getSuccessMessageResponse(reporterPlayer entity.Player, leaguePlayers []entity.LeaguePlayer, matchPlayers []entity.MatchPlayer) slack.MessageResponse {
	template := `
*%v* undoed a Match!

Standings *after* the undo:
%v
	`
	reporterDisplayName := reporterPlayer.DisplayName
	participantsText := getParticipantsText(leaguePlayers, matchPlayers)
	text := fmt.Sprintf(template, reporterDisplayName, participantsText)
	return slack.CreateChannelResponse(text)
}

func getParticipantsText(leaguePlayers []entity.LeaguePlayer, matchPlayers []entity.MatchPlayer) string {
	texts := []string{}

	for i := range matchPlayers {
		leaguePlayer := getParticipatingLeaguePlayer(matchPlayers[i], leaguePlayers)
		playerDisplayName := leaguePlayer.Player().DisplayName
		adjustedRating := leaguePlayer.Rating
		ratingChange := matchPlayers[i].RatingChange
		ratingChangeText := getRatingChangeText(adjustedRating, ratingChange)
		text := fmt.Sprintf("> *%v* was adjusted to *%v*. (%v)", playerDisplayName, adjustedRating, ratingChangeText)
		texts = append(texts, text)
	}

	return strings.Join(texts, "\n")
}

func getRatingChangeText(adjustedRating int, ratingChange int) string {
	originalRating := adjustedRating + ratingChange
	if originalRating > adjustedRating {
		return fmt.Sprintf("Lost *%v* rating", ratingChange)
	} else if originalRating < adjustedRating {
		return fmt.Sprintf("Gained *%v* rating", (-1)*ratingChange)
	}
	return "Unchanged"
}

func getParticipatingLeaguePlayer(matchPlayer entity.MatchPlayer, leaguePlayers []entity.LeaguePlayer) entity.LeaguePlayer {
	for i := range leaguePlayers {
		if matchPlayer.PlayerID == leaguePlayers[i].PlayerID {
			return leaguePlayers[i]
		}
	}
	return entity.LeaguePlayer{}
}

// NewOutputAdapterSlack factory
func NewOutputAdapterSlack() OutputAdapter {
	return &undoMatchOutputAdapterSlack{}
}
