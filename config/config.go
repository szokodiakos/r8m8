package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/szokodiakos/r8m8/logger"
)

// Setup config
func Setup() {
	err := godotenv.Load()
	if err != nil {
		logger.Get().Info(".env file was not found", err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Get().Fatal("Error reading in config.json", err)
	}
}
