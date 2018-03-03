package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/szokodiakos/r8m8/match"
)

func main() {
	setupConfig()

	connectionString := viper.GetString("mysql_connection_string")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("MySQL Database connect error", err)
	}
	matchRepository := match.NewRepository(db)
	matchService := match.NewService(matchRepository)
	matchSlackService := match.NewSlackService(matchService)

	e := echo.New()
	verificationToken := viper.GetString("slack_verification_token")
	match.NewSlackControllerHTTP(e, matchSlackService, verificationToken)

	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(port))
}

func setupConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("r8m8")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading in config.json", err)
	}
}
