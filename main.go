package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/league/leaderboard"
	"github.com/szokodiakos/r8m8/match/add"
	"github.com/szokodiakos/r8m8/match/undo"

	"github.com/szokodiakos/r8m8/rating"

	echoExtensions "github.com/szokodiakos/r8m8/echo"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
	"github.com/szokodiakos/r8m8/transaction"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/szokodiakos/r8m8/config"
	"github.com/szokodiakos/r8m8/match"
	sqlDB "github.com/szokodiakos/r8m8/sql"
)

func main() {
	config.Setup()

	sqlDialect := viper.GetString("sql_dialect")
	sqlConnectionString := viper.GetString("sql_connection_string")
	db, err := sqlx.Open(sqlDialect, sqlConnectionString)
	if err != nil {
		log.Fatal("Database connect error: ", err)
	}

	sqlDB.Execute(db, sqlDialect)

	database := sqlDB.NewSQLDB(db)
	transactionService := transaction.NewServiceSQL(database)

	playerRepository := entity.NewPlayerRepositorySQL()
	leagueRepository := entity.NewLeagueRepositorySQL()
	matchRepository := entity.NewMatchRepositorySQL()

	ratingStrategyElo := rating.NewStrategyElo()

	initialRating := 1500
	playerService := player.NewService(playerRepository)
	playerSlackService := player.NewSlackService()

	leaguePlayerService := league.NewPlayerService(playerService, playerRepository, initialRating)
	leagueService := league.NewService(playerService, leaguePlayerService, leagueRepository, playerRepository, initialRating)
	leagueSlackService := league.NewSlackService()

	verificationToken := viper.GetString("slack_verification_token")
	slackService := slack.NewService(verificationToken)

	e := echo.New()
	bodyParser := echoExtensions.BodyParser()
	slackTokenVerifier := slack.TokenVerifier(slackService)
	slackErrorHandler := slack.NewErrorHandler()
	slackHTTPErrorHandlerMiddleware := echoExtensions.ErrorHandlerMiddleware(slackErrorHandler)
	slackGroup := e.Group("/slack", bodyParser, slackTokenVerifier, slackHTTPErrorHandlerMiddleware)

	getLeaderboardInputAdapterSlack := leaderboard.NewGetLeaderboardInputAdapterSlack(slackService, leagueSlackService)
	getLeaderboardOutputAdapterSlack := leaderboard.NewGetLeaderboardOutputAdapterSlack()
	getLeaderboardUseCase := leaderboard.NewGetLeaderboardUseCase(transactionService, leagueRepository)
	leaderboard.NewGetLeaderboardControllerHTTP(slackGroup, getLeaderboardInputAdapterSlack, getLeaderboardOutputAdapterSlack, getLeaderboardUseCase)

	matchService := match.NewService(ratingStrategyElo, matchRepository, leaguePlayerService)

	addMatchInputAdapterSlack := add.NewAddMatchInputAdapterSlack(slackService, playerSlackService, leagueSlackService)
	addMatchOutputAdapterSlack := add.NewAddMatchOutputAdapterSlack()
	addMatchUseCase := add.NewAddMatchUseCase(transactionService, playerService, leagueService, leaguePlayerService, matchService, matchRepository, playerRepository, leagueRepository)
	add.NewAddMatchControllerHTTP(slackGroup, addMatchInputAdapterSlack, addMatchOutputAdapterSlack, addMatchUseCase)

	undoMatchInputAdapterSlack := undo.NewUndoMatchInputAdapterSlack(slackService, playerSlackService)
	undoMatchOutputAdapterSlack := undo.NewUndoMatchOutputAdapterSlack()
	undoMatchUseCase := undo.NewUndoMatchUseCase(transactionService, leaguePlayerService, matchRepository, leagueRepository)
	undo.NewUndoMatchControllerHTTP(slackGroup, undoMatchInputAdapterSlack, undoMatchOutputAdapterSlack, undoMatchUseCase)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}
