package match

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
)

// SlackService interface
type SlackService interface {
	AddMatch(values string) (slack.MessageResponse, error)
}

type matchSlackService struct {
	matchService Service
}

func (sms *matchSlackService) AddMatch(values string) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse
	rawPlayers := sms.parseValues(values)
	players, err := sms.parsePlayers(rawPlayers)
	if err != nil {
		return messageResponse, err
	}
	if sms.isPlayerCountUneven(players) {
		return messageResponse, errors.NewUnevenMatchPlayersError()
	}
	winnerPlayers := sms.getWinnerPlayers(players)
	loserPlayers := sms.getLoserPlayers(players)
	parsedMatch := Match{
		WinnerPlayers: winnerPlayers,
		LoserPlayers:  loserPlayers,
	}
	sms.matchService.AddMatch(parsedMatch)
	return messageResponse, nil
}

func (sms *matchSlackService) parseValues(values string) []string {
	parsedValues, _ := url.ParseQuery(values)
	text := parsedValues.Get("text")
	decodedText, _ := url.QueryUnescape(text)
	rawPlayers := strings.Split(decodedText, " ")
	return rawPlayers
}

func (sms *matchSlackService) parsePlayers(rawPlayers []string) ([]player.Player, error) {
	players := make([]player.Player, len(rawPlayers))
	for i, rawPlayer := range rawPlayers {
		player, err := sms.parsePlayer(rawPlayer)
		if err != nil {
			return nil, err
		}
		players[i] = player
	}
	return players, nil
}

func (sms *matchSlackService) parsePlayer(rawPlayer string) (player.Player, error) {
	var parsedPlayer player.Player
	playerRegexp, _ := regexp.Compile(`<@(.*)\|(.*)>`)
	results := playerRegexp.FindStringSubmatch(rawPlayer)
	if sms.isRawPlayerInvalid(results) {
		return parsedPlayer, errors.NewBadRawPlayerFormatError(rawPlayer)
	}
	playerID := results[1]
	playerName := results[2]
	parsedPlayer = player.Player{
		ID:   playerID,
		Name: playerName,
	}
	return parsedPlayer, nil
}

func (sms *matchSlackService) isRawPlayerInvalid(results []string) bool {
	if len(results) != 3 {
		return true
	}
	return false
}

func (sms *matchSlackService) isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func (sms *matchSlackService) getWinnerPlayers(players []player.Player) []player.Player {
	lowerhalfPlayers := players[:(len(players) / 2)]
	return lowerhalfPlayers
}

func (sms *matchSlackService) getLoserPlayers(players []player.Player) []player.Player {
	upperhalfPlayers := players[(len(players) / 2):]
	return upperhalfPlayers
}

// NewSlackService factory
func NewSlackService(matchService Service) SlackService {
	return &matchSlackService{
		matchService: matchService,
	}
}
