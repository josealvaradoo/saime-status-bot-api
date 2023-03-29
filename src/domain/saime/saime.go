package saime

import (
	"fmt"
	"net/http"

	"github.com/josealvaradoo/saime-status-bot/src/domain/telegram"
	"github.com/josealvaradoo/saime-status-bot/src/model"
	repo "github.com/josealvaradoo/saime-status-bot/src/repository/saime"
)

func prepareBotMessage(status string) string {
	if status == model.Offline {
		return fmt.Sprintf("❌ La página del SAIME está: %s", status)
	}
	return fmt.Sprintf("✅ La página del SAIME está: %s", status)
}

func sendTelegramMessage(bot *telegram.Bot, prevStatus string, status string) {
	if prevStatus != status {
		message := prepareBotMessage(status)

		err := bot.Notify(message)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(message)
		}
	}
}

func Get() (model.Saime, error) {
	cache, err := repo.Get()

	if err != nil {
		return model.Saime{}, err
	}

	return model.Saime{Status: cache.Status}, nil
}

func Post() (string, error) {
	bot := telegram.Bot{}
	var status string = model.Online
	err := CheckAvailability()

	if err != nil {
		status = model.Offline
	}

	prev, err := Get()

	if err != nil {
		prev.Status = ""
	}

	sendTelegramMessage(&bot, prev.Status, status)

	r, err := repo.Update(status)

	if err != nil {
		return r.Status, err
	}

	return r.Status, nil
}

func CheckAvailability() error {
	url := "https://siic.saime.gob.ve"
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Hubo un error intetando hacer get a " + url)
	}

	return err
}
