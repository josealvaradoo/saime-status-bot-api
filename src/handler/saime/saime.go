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

func prepareBotMessage(status string) string {
	if status == model.Offline {
		return fmt.Sprintf("❌ La pagian del SAIME está: %s", status)
	}
	return fmt.Sprintf("✅ La pagian del SAIME está: %s", status)
}

func sendTelegramMessage(bot *telegram.Bot, prevStatus string, status string) {
	if prevStatus != status {
		message := prepareBotMessage(status)

		err := bot.Notify(message)
		if err != nil {
			fmt.Println(err)
		}
	}
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
	var status string = model.Online
	err := saime.CheckAvailability()

	if err != nil {
		fmt.Println("Error checking availability")
		fmt.Println(err.Error())
		status = model.Offline
	}

	prev, err := saime.Get()

	if err != nil {
		sendTelegramMessage(&bot, "", status)
	}

	sendTelegramMessage(&bot, prev.Status, status)

	err = saime.Set(status)

	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadGateway, err.Error()))
	}

	return c.JSON(model.Saime{Status: status})
}
