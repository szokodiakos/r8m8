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
	matchRepository := match.NewRepositorySQL()
	ratingStrategyElo := rating.NewStrategyElo()
	ratingRepository := rating.NewRepositorySQL()
	detailsRepository := details.NewRepositorySQL()
	ratingService := rating.NewService(ratingStrategyElo, ratingRepository, detailsRepository)
	playerRepository := player.NewRepository()
	initialRating := 1500
	playerService := player.NewService(playerRepository, ratingRepository, initialRating)
	leagueRepository := league.NewRepositorySQL()
	leagueService := league.NewService(leagueRepository)
	matchService := match.NewService(matchRepository, ratingService, playerService, leagueService)
	verificationToken := viper.GetString("slack_verification_token")
	slackService := slack.NewService(verificationToken)
	playerSlackService := player.NewSlackService()
	transactionService := transaction.NewServiceSQL(database)
	leagueSlackService := league.NewSlackService()
	matchSlackService := match.NewSlackService(matchService, slackService, playerSlackService, leagueSlackService, transactionService)
	e := echo.New()
	bodyParser := echoExtensions.BodyParser()
	slackTokenVerifier := slack.TokenVerifier(slackService)
	slackErrorHandler := slack.NewErrorHandler()
	httpErrorHandler := echoExtensions.ErrorHandlerMiddleware(slackErrorHandler)
	slackGroup := e.Group("/slack", bodyParser, slackTokenVerifier, httpErrorHandler)
	match.NewSlackControllerHTTP(slackGroup, matchSlackService, slackService)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}
