package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Setup config
func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetEnvPrefix("r8m8")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading in config.json: ", err)
	}
}
