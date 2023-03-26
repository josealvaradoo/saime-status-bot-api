package saime

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/domain/saime"
	"github.com/josealvaradoo/saime-status-bot/src/domain/telegram"
	"github.com/josealvaradoo/saime-status-bot/src/model"
)

type Handler interface {
	Check(c *fiber.Ctx) error
	Set(c *fiber.Ctx, key string) error
}

func Get(c *fiber.Ctx) error {
	value, err := saime.Get()

	if err != nil {
		c.JSON(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	return c.JSON(value)
}

func Post(c *fiber.Ctx) error {
	bot := telegram.Bot{}
	var message string
	var status string = model.Online
	err := saime.CheckAvailability()

	if err != nil {
		status = model.Offline
	}

	err = saime.Set(status)

	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadGateway, err.Error()))
	}

	prev, err := saime.Get()

	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	if prev.Status != status {
		if status == model.Offline {
			message = fmt.Sprintf("❌ La pagian del SAIME está: %s", status)
		} else {
			message = fmt.Sprintf("✅ La pagian del SAIME está: %s", status)
		}

		err = bot.Notify(message)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.JSON(model.Saime{Status: status})
}
