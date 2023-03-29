package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/josealvaradoo/saime-status-bot/src/handler/saime"
)

func Saime(app *fiber.App) {
	api := app.Group("/api/v1")
	status := api.Group("/status")

	status.Get("", handler.Get)
	status.Post("", handler.Post)
}
