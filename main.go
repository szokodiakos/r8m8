package main

import (
	"database/sql"
	"log"

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

	matchRepository := match.NewRepository(db)
	matchService := match.NewService(matchRepository)
	matchSlackService := match.NewSlackService(matchService)

	e := echo.New()
	verificationToken := viper.GetString("slack_verification_token")
	match.NewSlackControllerHTTP(e, matchSlackService, verificationToken)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}
