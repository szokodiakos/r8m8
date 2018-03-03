package slack

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/szokodiakos/r8m8/match"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack/errors"
)

// MatchService Slack
type MatchService interface {
	ParseMatch(values string) (match.Match, error)
}

type matchService struct {
}

func (ms *matchService) ParseMatch(values string) (match.Match, error) {
	var parsedMatch match.Match
	rawPlayers := ms.parseValues(values)
	players, err := ms.parsePlayers(rawPlayers)
	if err != nil {
		return parsedMatch, err
	}
	if ms.isPlayerCountUneven(players) {
		return parsedMatch, errors.NewUnevenMatchPlayersError()
	}
	winnerPlayers := ms.getWinnerPlayers(players)
	loserPlayers := ms.getLoserPlayers(players)
	parsedMatch = match.Match{
		WinnerPlayers: winnerPlayers,
		LoserPlayers:  loserPlayers,
	}
	return parsedMatch, nil
}

func (ms *matchService) parseValues(values string) []string {
	parsedValues, _ := url.ParseQuery(values)
	text := parsedValues.Get("text")
	decodedText, _ := url.QueryUnescape(text)
	rawPlayers := strings.Split(decodedText, " ")
	return rawPlayers
}

func (ms *matchService) parsePlayers(rawPlayers []string) ([]player.Player, error) {
	players := make([]player.Player, len(rawPlayers))
	for i, rawPlayer := range rawPlayers {
		player, err := ms.parsePlayer(rawPlayer)
		if err != nil {
			return nil, err
		}
		players[i] = player
	}
	return players, nil
}

func (ms *matchService) parsePlayer(rawPlayer string) (player.Player, error) {
	var parsedPlayer player.Player
	playerRegexp, _ := regexp.Compile(`<@(.*)\|(.*)>`)
	results := playerRegexp.FindStringSubmatch(rawPlayer)
	if ms.isRawPlayerInvalid(results) {
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

func (ms *matchService) isRawPlayerInvalid(results []string) bool {
	if len(results) != 3 {
		return true
	}
	return false
}

func (ms *matchService) isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func (ms *matchService) getWinnerPlayers(players []player.Player) []player.Player {
	lowerhalfPlayers := players[:(len(players) / 2)]
	return lowerhalfPlayers
}

func (ms *matchService) getLoserPlayers(players []player.Player) []player.Player {
	upperhalfPlayers := players[(len(players) / 2):]
	return upperhalfPlayers
}

// NewMatchService creates a service
func NewMatchService() MatchService {
	return &matchService{}
}
