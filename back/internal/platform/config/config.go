package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	Port string
}

var config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	err := godotenv.Load("configs/server.env")
	if err != nil {
		log.Println("Info: No .env file found, using system environment variables.")
	}
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
}

func GetConfig() ConfigList {
	return config
}
