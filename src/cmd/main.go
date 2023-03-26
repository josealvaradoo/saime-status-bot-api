package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/domain/telegram"
	handler "github.com/josealvaradoo/saime-status-bot/src/handler/saime"
	storage "github.com/josealvaradoo/saime-status-bot/src/storage/cache"
	"github.com/josealvaradoo/saime-status-bot/src/utils"
)

func main() {
	bot := telegram.Bot{}
	driver := storage.Redis
	storage.New(driver)
	app := fiber.New()

	api := app.Group("/api/v1")

	status := api.Group("/status")

	status.Get("", handler.Get)
	status.Post("", handler.Post)

	go bot.New()
	log.Fatal(app.Listen(fmt.Sprintf(":%s", utils.Env(utils.PORT))))
}
