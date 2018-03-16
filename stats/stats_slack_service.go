package stats

import (
	"fmt"
	"strings"

	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/slack"
	"github.com/szokodiakos/r8m8/transaction"
)

// SlackService interface
type SlackService interface {
	GetLeaderboard(values string) (slack.MessageResponse, error)
}

type statsSlackService struct {
	transactionService transaction.Service
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

	transaction, err := s.transactionService.Start()
	if err != nil {
		return messageResponse, err
	}

	teamID := requestValues.TeamID
	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := s.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)
	leaderboard, err := s.statsService.GetLeaderboard(transaction, league)
	if err != nil {
		return messageResponse, s.transactionService.Rollback(transaction, err)
	}

	err = s.transactionService.Commit(transaction)

	messageResponse = getSuccessMessageResponse(leaderboard)
	return messageResponse, nil
}

func getSuccessMessageResponse(leaderboard Leaderboard) slack.MessageResponse {
	text := fmt.Sprintf(`
:fire: TOP 10 Leaderboard for *%v* :fire:

%v
	`, leaderboard.DisplayName, getPlayerStatsTexts(leaderboard.Players, 10))
	return slack.CreateChannelResponse(text)
}

func getPlayerStatsTexts(playersStats []PlayerStats, c int) string {
	count := len(playersStats)
	if c < count {
		count = c
	}

	playerText := make([]string, count)
	for i := 0; i < count; i++ {
		icon := getIcon(i + 1)
		displayName := playersStats[i].DisplayName
		rating := playersStats[i].Rating
		winCount := playersStats[i].WinCount
		matchCount := playersStats[i].MatchCount
		lossCount := matchCount - winCount
		playerText[i] = fmt.Sprintf("> *%v*	%v	*%v*	(%v Win / %v Loss)", icon, displayName, rating, winCount, lossCount)
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

// NewSlackService factory
func NewSlackService(statsService Service, leagueSlackService league.SlackService, slackService slack.Service, transactionService transaction.Service) SlackService {
	return &statsSlackService{
		statsService:       statsService,
		leagueSlackService: leagueSlackService,
		slackService:       slackService,
		transactionService: transactionService,
	}
}
