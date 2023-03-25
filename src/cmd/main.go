package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/handler"
	"github.com/josealvaradoo/saime-status-bot/src/storage"
)

func main() {
	driver := storage.Redis
	storage.New(driver)
	app := fiber.New()

	api := app.Group("/api/v1")

	status := api.Group("/status")

	status.Get("", handler.Get)
	status.Post("", handler.Post)

	log.Fatal(app.Listen(":3000"))
}
