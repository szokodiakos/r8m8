package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/szokodiakos/r8m8/league"

	"github.com/szokodiakos/r8m8/details"

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

	detailsRepository := details.NewRepositorySQL()
	ratingRepository := rating.NewRepositorySQL()
	playerRepository := player.NewRepositorySQL()
	leaguePlayerRepository := league.NewPlayerRepositorySQL()
	leagueRepository := league.NewRepositorySQL(leaguePlayerRepository)
	matchRepository := match.NewRepositorySQL()
	matchPlayerRepository := match.NewPlayerRepositorySQL()

	ratingStrategyElo := rating.NewStrategyElo()
	ratingService := rating.NewService(ratingStrategyElo, ratingRepository, detailsRepository)

	initialRating := 1500
	playerService := player.NewService(playerRepository, ratingRepository, initialRating)
	playerSlackService := player.NewSlackService()

	leagueService := league.NewService(leagueRepository)
	leagueSlackService := league.NewSlackService()

	verificationToken := viper.GetString("slack_verification_token")
	slackService := slack.NewService(verificationToken)

	e := echo.New()
	bodyParser := echoExtensions.BodyParser()
	slackTokenVerifier := slack.TokenVerifier(slackService)
	slackErrorHandler := slack.NewErrorHandler()
	slackHTTPErrorHandlerMiddleware := echoExtensions.ErrorHandlerMiddleware(slackErrorHandler)
	slackGroup := e.Group("/slack", bodyParser, slackTokenVerifier, slackHTTPErrorHandlerMiddleware)

	getLeaderboardInputAdapterSlack := league.NewGetLeaderboardInputAdapterSlack(slackService, leagueSlackService)
	getLeaderboardOutputAdapterSlack := league.NewGetLeaderboardOutputAdapterSlack()
	getLeaderboardUseCase := league.NewGetLeaderboardUseCase(transactionService, leagueRepository)
	league.NewGetLeaderboardControllerHTTP(slackGroup, getLeaderboardInputAdapterSlack, getLeaderboardOutputAdapterSlack, getLeaderboardUseCase)

	matchService := match.NewService(matchRepository, matchPlayerRepository)
	addMatchInputAdapterSlack := match.NewAddMatchInputAdapterSlack(slackService, playerSlackService, leagueSlackService)
	addMatchOutputAdapterSlack := match.NewAddMatchOutputAdapterSlack()
	addMatchUseCase := match.NewAddMatchUseCase(transactionService, matchRepository, ratingService, playerService, leagueService, matchService)
	match.NewAddMatchControllerHTTP(slackGroup, addMatchInputAdapterSlack, addMatchOutputAdapterSlack, addMatchUseCase)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}
