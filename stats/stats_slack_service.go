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
		s.transactionService.Rollback(transaction)
		return messageResponse, err
	}

	err = s.transactionService.Commit(transaction)

	messageResponse = getSuccessMessageResponse(leaderboard)
	return messageResponse, nil
}

func getSuccessMessageResponse(leaderboard Leaderboard) slack.MessageResponse {
	text := fmt.Sprintf(`
Leaderboard for %v: *Player* @ *Rating* _(win/loss)_
%v
	`, leaderboard.DisplayName, getLeaderboardPlayersText(leaderboard.Players))
	return slack.MessageResponse{
		Text: text,
	}
}

func getLeaderboardPlayersText(leaderboardPlayers []LeaderboardPlayer) string {
	playerText := make([]string, len(leaderboardPlayers))
	for i := range leaderboardPlayers {
		displayName := leaderboardPlayers[i].DisplayName
		rating := leaderboardPlayers[i].Rating
		winCount := leaderboardPlayers[i].WinCount
		matchCount := leaderboardPlayers[i].MatchCount
		lossCount := matchCount - winCount
		icon := getIcon(i + 1)
		playerText[i] = fmt.Sprintf("> *%v* %v *%v* (%v/%v)", displayName, icon, rating, winCount, lossCount)
	}
	return strings.Join(playerText, "\n")
}

func getIcon(place int) string {
	if place == 1 {
		return ":first_place_medal:"
	}
	if place == 2 {
		return ":second_place_medal:"
	}
	if place == 3 {
		return ":third_place_medal:"
	}
	return "@"
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
