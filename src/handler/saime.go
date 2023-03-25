package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/domain"
	"github.com/josealvaradoo/saime-status-bot/src/model"
)

type Handler interface {
	Check(c *fiber.Ctx) error
	Set(c *fiber.Ctx, key string) error
}

func Get(c *fiber.Ctx) error {
	value, err := domain.Get()

	if err != nil {
		c.JSON(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	return c.JSON(value)
}

func Post(c *fiber.Ctx) error {
	var status string = model.Online
	err := domain.CheckAvailability()

	if err != nil {
		status = model.Offline
	}

	err = domain.Set(status)

	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadGateway, err.Error()))
	}

	return c.JSON(model.Saime{Status: status})
}
