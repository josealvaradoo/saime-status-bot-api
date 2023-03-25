package domain

import (
	"os/exec"

	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/repository"
)

func Get() (model.Saime, error) {
	cache, err := repository.Get()

	if err != nil {
		return model.Saime{}, err
	}

	return model.Saime{Status: cache.Status}, nil
}

func Set(status string) error {
	_, err := repository.Update(status)

	if err != nil {
		return err
	}

	return nil
}

func CheckAvailability() error {
	url := "https://siic.saime.gob.ve"
	curl := exec.Command("curl", url)

	_, err := curl.Output()
	if err != nil {
		return err
	} else {
		return nil
	}
}
