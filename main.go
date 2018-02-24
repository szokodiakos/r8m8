package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/szokodiakos/r8m8/leaderboard"
)

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
		log.Fatal("Error reading in config.json")
	}
}

func main() {
	setupConfig()

	lb := leaderboard.NewService(leaderboard.NewSlackPrinter(viper.GetString("webhook")))
	lb.GetLeaderboard()
}
