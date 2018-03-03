package main

import (
	"database/sql"
	"log"

	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"

	"github.com/szokodiakos/r8m8/transaction"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/szokodiakos/r8m8/config"
	"github.com/szokodiakos/r8m8/match"
	sqlMigrate "github.com/szokodiakos/r8m8/sql"
)

func main() {
	config.Setup()

	sqlDialect := viper.GetString("sql_dialect")
	sqlConnectionString := viper.GetString("sql_connection_string")
	db, err := sql.Open(sqlDialect, sqlConnectionString)
	if err != nil {
		log.Fatal("Database connect error: ", err)
	}

	sqlMigrate.Execute(db, sqlDialect)

	matchRepository := match.NewRepositorySQL(db)
	transactionService := transaction.NewServiceSQL(db)
	matchService := match.NewService(transactionService, matchRepository)
	slackService := slack.NewService()
	playerSlackRepository := player.NewSlackRepository(db)
	playerRepository := player.NewRepository(db)
	playerService := player.NewService(playerRepository)
	playerSlackParserService := player.NewSlackParserService()
	playerSlackService := player.NewSlackService(playerSlackRepository, transactionService, playerService, playerSlackParserService)
	matchSlackService := match.NewSlackService(matchService, slackService, playerSlackService)

	e := echo.New()
	verificationToken := viper.GetString("slack_verification_token")
	match.NewSlackControllerHTTP(e, matchSlackService, slackService, verificationToken)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}
