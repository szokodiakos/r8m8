package stats

import (
	"fmt"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/slack"
)

type statsOutputAdapterSlack struct {
}

func (g *statsOutputAdapterSlack) Handle(output Output, err error) (interface{}, error) {
	if err != nil {
		return getErrorMessageResponse(err)
	}

	leaguePlayer := output.LeaguePlayer
	return getSuccessMessageResponse(leaguePlayer), nil
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

func getSuccessMessageResponse(leaguePlayer entity.LeaguePlayer) slack.MessageResponse {
	text := fmt.Sprintf(`Hello`)
	return slack.CreateChannelResponse(text)
}

// NewOutputAdapterSlack factory
func NewOutputAdapterSlack() OutputAdapter {
	return &statsOutputAdapterSlack{}
}
