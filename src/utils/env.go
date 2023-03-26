package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	PORT               = "PORT"
	REDIS_HOST         = "REDIS_HOST"
	REDIS_PORT         = "REDIS_PORT"
	TELEGRAM_BOT_TOKEN = "TELEGRAM_BOT_TOKEN"
	TELEGRAM_CHAT_ID   = "TELEGRAM_CHAT_ID"
)

func Env(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")

		err := godotenv.Load("../../.env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	return os.Getenv(key)
}
