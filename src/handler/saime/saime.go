package saime

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/domain/saime"
	"github.com/josealvaradoo/saime-status-bot/src/model"
)

type Handler interface {
	Check(c *fiber.Ctx) error
	Set(c *fiber.Ctx, key string) error
}

func Get(c *fiber.Ctx) error {
	value, err := saime.Get()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	return c.JSON(value)
}

func Post(c *fiber.Ctx) error {
	status, err := saime.Post()

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.NewError(fiber.StatusBadGateway, err.Error()))
	}

	return c.JSON(model.Saime{Status: status})
}
