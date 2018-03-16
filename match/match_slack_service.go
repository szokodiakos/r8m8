package match

import (
	"fmt"
	"strings"

	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
	"github.com/szokodiakos/r8m8/stats"
	statsModel "github.com/szokodiakos/r8m8/stats/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// SlackService interface
type SlackService interface {
	Add(values string) (slack.MessageResponse, error)
}

type matchSlackService struct {
	matchService       Service
	slackService       slack.Service
	playerSlackService player.SlackService
	leagueSlackService league.SlackService
	transactionService transaction.Service
	statsService       stats.Service
}

func (m *matchSlackService) Add(values string) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse

	requestValues, err := m.slackService.ParseRequestValues(values)
	if err != nil {
		return messageResponse, err
	}

	text := requestValues.Text
	teamID := requestValues.TeamID
	players, err := m.playerSlackService.ToPlayers(text, teamID)
	if err != nil {
		return messageResponse, err
	}

	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := m.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := m.playerSlackService.ToPlayer(teamID, userID, userName)
	tr, err := m.transactionService.Start()
	if err != nil {
		return messageResponse, err
	}

	matchID, err := m.matchService.Add(tr, players, league, reporterPlayer)
	if err != nil {
		return messageResponse, m.transactionService.Rollback(tr, err)
	}

	matchStats, err := m.statsService.GetMatchStats(tr, matchID)
	if err != nil {
		return messageResponse, m.transactionService.Rollback(tr, err)
	}

	err = m.transactionService.Commit(tr)

	messageResponse = getSuccessMessageResponse(matchStats)
	return messageResponse, err
}

func getSuccessMessageResponse(matchStats statsModel.MatchStats) slack.MessageResponse {
	template := `
*%v* recorded a new Match! Good Game! :muscle:

*Winners* :trophy:
%v
*Losers*
%v
	`
	reporterDisplayName := matchStats.ReporterPlayer.DisplayName
	winnersDetailsText := getMatchPlayerStatsTexts(matchStats.WinnerMatchPlayersStats)
	losersDetailsText := getMatchPlayerStatsTexts(matchStats.LoserMatchPlayersStats)
	text := fmt.Sprintf(template, reporterDisplayName, winnersDetailsText, losersDetailsText)
	return slack.CreateChannelResponse(text)
}

func getMatchPlayerStatsTexts(matchPlayersStats []statsModel.MatchPlayerStats) string {
	texts := []string{}
	for i := range matchPlayersStats {
		displayName := matchPlayersStats[i].Player.DisplayName
		ratingChange := matchPlayersStats[i].Details.RatingChange
		rating := matchPlayersStats[i].Rating.Rating
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

// NewSlackService factory
func NewSlackService(matchService Service, slackService slack.Service, playerSlackService player.SlackService, leagueSlackService league.SlackService, transactionService transaction.Service, statsService stats.Service) SlackService {
	return &matchSlackService{
		matchService:       matchService,
		slackService:       slackService,
		playerSlackService: playerSlackService,
		leagueSlackService: leagueSlackService,
		transactionService: transactionService,
		statsService:       statsService,
	}
}
