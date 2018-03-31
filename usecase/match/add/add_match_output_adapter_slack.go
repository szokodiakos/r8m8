package add

import (
	"fmt"
	"strings"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/match/errors"
	playerErrors "github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/slack"
)

type addMatchOutputAdapterSlack struct {
}

func (a *addMatchOutputAdapterSlack) Handle(output Output, err error) (interface{}, error) {
	match := output.Match

	if err != nil {
		return getErrorMessageResponse(err)
	}

	return getSuccessMessageResponse(match), nil
}

func getErrorMessageResponse(err error) (slack.MessageResponse, error) {
	switch err.(type) {
	case *errors.ReporterPlayerNotInLeagueError:
		return getReporterPlayerNotInLeagueResponse(), nil
	case *playerErrors.DuplicatedPlayerExistsError:
		return getDuplicatedPlayerExistsResponse(), nil
	case *playerErrors.UnevenMatchPlayersError:
		return getUnevenMatchPlayersResponse(), nil
	case *playerErrors.BadSlackPlayerFormatError:
		badSlackPlayerFormatError := err.(*playerErrors.BadSlackPlayerFormatError)
		return getBadSlackPlayerFormatResponse(badSlackPlayerFormatError), nil
	default:
		return slack.MessageResponse{}, err
	}
}

func getReporterPlayerNotInLeagueResponse() slack.MessageResponse {
	text := `
> Darn! You must be the participant of at least one match (including this one). :hushed:
> :exclamation: Please play a match before posting! :exclamation:
`
	return slack.CreateDirectResponse(text)
}

func getDuplicatedPlayerExistsResponse() slack.MessageResponse {
	text := `
> Darn! You provided a player more than once. :hushed:
`
	return slack.CreateDirectResponse(text)
}

func getUnevenMatchPlayersResponse() slack.MessageResponse {
	text := `
> Darn! Reported players are uneven! :hushed:
> :exclamation: Make sure you report even number of players! :exclamation:
`
	return slack.CreateDirectResponse(text)
}

func getBadSlackPlayerFormatResponse(err *playerErrors.BadSlackPlayerFormatError) slack.MessageResponse {
	text := fmt.Sprintf(`
> Darn! Reported player "%v" is not valid! :hushed:
`, err.SlackPlayer)

	return slack.CreateDirectResponse(text)
}

func getSuccessMessageResponse(match entity.Match) slack.MessageResponse {
	template := `
*%v* recorded a new Match! Good Game! :muscle:

*Winners* :trophy:
%v
*Losers*
%v
	`
	reporterDisplayName := match.ReporterPlayer().DisplayName
	winnerMatchPlayersText := getMatchPlayersText(match.WinnerMatchPlayers())
	loserMatchPlayersText := getMatchPlayersText(match.LoserMatchPlayers())
	text := fmt.Sprintf(template, reporterDisplayName, winnerMatchPlayersText, loserMatchPlayersText)
	return slack.CreateChannelResponse(text)
}

func getMatchPlayersText(matchPlayers []entity.MatchPlayer) string {
	texts := []string{}
	for _, matchPlayer := range matchPlayers {
		leaguePlayer := matchPlayer.LeaguePlayer()
		player := leaguePlayer.Player()

		displayName := player.DisplayName
		ratingChange := matchPlayer.RatingChange
		rating := leaguePlayer.Rating
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
func NewAddMatchOutputAdapterSlack() OutputAdapter {
	return &addMatchOutputAdapterSlack{}
}
