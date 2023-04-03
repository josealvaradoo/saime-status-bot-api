package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/domain/telegram"
	"github.com/josealvaradoo/saime-status-bot/src/router"
	"github.com/josealvaradoo/saime-status-bot/src/storage/firestore"

	"github.com/josealvaradoo/saime-status-bot/src/utils"
)

func main() {
	// Initialize Firestore
	firestore.New()

	// Define Fiber API
	app := fiber.New()
	router.Saime(app)

	// Initialize Telegram bot
	bot := telegram.Bot{}
	go bot.New()

	// Initialize API
	log.Fatal(app.Listen(fmt.Sprintf(":%s", utils.Env(utils.PORT))))
}
