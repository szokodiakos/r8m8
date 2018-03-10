package main

import (
	"database/sql"
	"log"

	"github.com/szokodiakos/r8m8/rating"

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
	db, err := sql.Open(sqlDialect, sqlConnectionString)
	if err != nil {
		log.Fatal("Database connect error: ", err)
	}

	sqlDB.Execute(db, sqlDialect)

	database := sqlDB.NewSQLDB(db)
	matchRepository := match.NewRepositorySQL()
	ratingService := rating.NewService()
	matchDetailsRepository := match.NewDetailsRepositorySQL()
	matchDetailsService := match.NewDetailsService(matchDetailsRepository)
	playerRepository := player.NewRepository()
	playerService := player.NewService(playerRepository)
	matchService := match.NewService(matchRepository, ratingService, playerService, matchDetailsService)
	slackService := slack.NewService()
	playerSlackRepository := player.NewSlackRepository()
	playerSlackParserService := player.NewSlackParserService()
	playerSlackService := player.NewSlackService(playerSlackRepository, playerService, playerSlackParserService)
	transactionService := transaction.NewServiceSQL(database)
	matchSlackService := match.NewSlackService(matchService, slackService, playerSlackService, transactionService)

	e := echo.New()
	verificationToken := viper.GetString("slack_verification_token")
	match.NewSlackControllerHTTP(e, matchSlackService, slackService, verificationToken)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}
