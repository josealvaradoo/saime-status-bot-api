package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/domain/saime"
	"github.com/josealvaradoo/saime-status-bot/src/domain/telegram"
	"github.com/josealvaradoo/saime-status-bot/src/router"

	"github.com/josealvaradoo/saime-status-bot/src/utils"
	"github.com/robfig/cron/v3"
)

func main() {
	// Define Fiber API
	app := fiber.New()
	router.Saime(app)

	// Initialize Cron job
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		saime.Post()
	})
	c.Start()

	// Initialize Telegram bot
	bot := telegram.Bot{}
	go bot.New()

	// Initialize API
	log.Fatal(app.Listen(fmt.Sprintf(":%s", utils.Env(utils.PORT))))
}
