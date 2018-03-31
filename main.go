package main

import (
	"database/sql"
	"fmt"

	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/logger"
	"github.com/szokodiakos/r8m8/usecase/leaderboard"
	"github.com/szokodiakos/r8m8/usecase/match/undo"

	"github.com/szokodiakos/r8m8/rating"

	echoExtensions "github.com/szokodiakos/r8m8/echo"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
	"github.com/szokodiakos/r8m8/transaction"
	"github.com/szokodiakos/r8m8/usecase/match/add"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/szokodiakos/r8m8/config"
	"github.com/szokodiakos/r8m8/match"
	sqlDB "github.com/szokodiakos/r8m8/sql"
)

func main() {
	config.Setup()

	sentryDSN := viper.GetString("sentry_dsn")
	logFormat := viper.GetString("log_format")
	logger.Setup(sentryDSN, logFormat)

	sqlDialect := viper.GetString("sql_dialect")
	sqlConnectionString := viper.GetString("sql_connection_string")
	db, err := sql.Open(sqlDialect, sqlConnectionString)
	if err != nil {
		logger.Get().Fatal("Database connect error", err)
	}

	sqlDB.Execute(db, sqlDialect)

	database := sqlDB.NewSQLDB(db, sqlDialect)
	transactionService := transaction.NewServiceSQL(database)

	playerRepository := player.NewPlayerRepositorySQL()
	leagueRepository := league.NewLeagueRepositorySQL()
	matchRepository := match.NewMatchRepositorySQL()

	ratingStrategyElo := rating.NewStrategyElo()

	initialRating := 1500
	playerService := player.NewService(playerRepository)
	playerSlackService := player.NewSlackService()

	leaguePlayerService := league.NewPlayerService(playerService, playerRepository, initialRating)
	leagueService := league.NewService(leagueRepository, playerRepository, initialRating)
	leagueSlackService := league.NewSlackService()

	verificationToken := viper.GetString("slack_verification_token")
	slackService := slack.NewService(verificationToken)

	server := echo.New()
	server.HideBanner = true
	bodyParser := echoExtensions.BodyParser()
	slackTokenVerifier := slack.TokenVerifier(slackService)
	slackErrorHandler := slack.NewErrorHandler()
	slackHTTPErrorHandlerMiddleware := echoExtensions.ErrorHandlerMiddleware(slackErrorHandler)
	slackGroup := server.Group("/slack", bodyParser, slackTokenVerifier, slackHTTPErrorHandlerMiddleware)

	leaderboardInputAdapterSlack := leaderboard.NewInputAdapterSlack(slackService, leagueSlackService)
	leaderboardOutputAdapterSlack := leaderboard.NewOutputAdapterSlack()
	leaderboardUseCase := leaderboard.NewUseCase(transactionService, leagueRepository)
	leaderboard.NewControllerHTTP(slackGroup, leaderboardInputAdapterSlack, leaderboardOutputAdapterSlack, leaderboardUseCase)

	matchService := match.NewService(ratingStrategyElo)

	addMatchInputAdapterSlack := add.NewInputAdapterSlack(slackService, playerSlackService, leagueSlackService)
	addMatchOutputAdapterSlack := add.NewOutputAdapterSlack()
	addMatchUseCase := add.NewUseCase(transactionService, playerService, leagueService, leaguePlayerService, matchService, matchRepository, playerRepository, leagueRepository)
	add.NewControllerHTTP(slackGroup, addMatchInputAdapterSlack, addMatchOutputAdapterSlack, addMatchUseCase)

	undoMatchInputAdapterSlack := undo.NewInputAdapterSlack(slackService, playerSlackService)
	undoMatchOutputAdapterSlack := undo.NewOutputAdapterSlack()
	undoMatchUseCase := undo.NewUseCase(transactionService, leaguePlayerService, matchRepository, leagueRepository)
	undo.NewControllerHTTP(slackGroup, undoMatchInputAdapterSlack, undoMatchOutputAdapterSlack, undoMatchUseCase)

	port := viper.GetString("port")
	logger.Get().Infof("Server starting on port %v.", port)
	logger.Get().Fatal("Server could not start", server.Start(fmt.Sprintf(":%v", port)))
}
