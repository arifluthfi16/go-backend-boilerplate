package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
	Will load config from .env.example
*/

func LoadConfig () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files, make sure .env file exists in project root")
	}

	// Application Config
	ServerConfig.AppConfig.AppName = os.Getenv("APP_NAME")
	ServerConfig.AppConfig.AppVersion = os.Getenv("APP_VERSION")

	// Database Config
	ServerConfig.DBConfig.DBPort = os.Getenv("DB_PORT")
	ServerConfig.DBConfig.DBName = os.Getenv("DB_NAME")
	ServerConfig.DBConfig.DBUser = os.Getenv("DB_USERNAME")
	ServerConfig.DBConfig.DBHost = os.Getenv("DB_HOST")
	ServerConfig.DBConfig.DBPass = os.Getenv("DB_PASSWORD")
}