package telegram

import (
	"log"
	"strconv"

	telegramAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/josealvaradoo/saime-status-bot/src/utils"
)

type Bot struct {
	Bot *telegramAPI.BotAPI
}

var (
	Telegram = &Bot{}
)

func (t *Bot) New() {
	var err error

	Telegram.Bot, err = telegramAPI.NewBotAPI(utils.Env(utils.TELEGRAM_BOT_TOKEN))

	if err != nil {
		log.Fatal(err)
	}

	Telegram.Bot.Debug = true

	log.Printf("Authorized on account %s", Telegram.Bot.Self.UserName)
}

func (t *Bot) Notify(message string) error {
	channelID, err := strconv.ParseInt(utils.Env(utils.TELEGRAM_CHAT_ID), 0, 64)

	if err != nil {
		log.Fatal(err)
	}

	msg := telegramAPI.NewMessage(channelID, message)
	_, err = Telegram.Bot.Send(msg)

	return err
}
