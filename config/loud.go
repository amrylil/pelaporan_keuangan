package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}

func LoadDBConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}

	var res = new(DatabaseConfig)

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DB_USER = val
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DB_PASS = val
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DB_HOST = val
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		res.DB_PORT = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DB_NAME = val
	}

	return res
}
