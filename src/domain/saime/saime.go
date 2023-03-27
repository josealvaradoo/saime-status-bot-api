package saime

import (
	"fmt"
	"os/exec"

	"github.com/josealvaradoo/saime-status-bot/src/model"
	repo "github.com/josealvaradoo/saime-status-bot/src/repository/saime"
)

func Get() (model.Saime, error) {
	cache, err := repo.Get()

	if err != nil {
		return model.Saime{}, err
	}

	return model.Saime{Status: cache.Status}, nil
}

func Set(status string) error {
	_, err := repo.Update(status)

	if err != nil {
		return err
	}

	return nil
}

func CheckAvailability() error {
	url := "https://siic.saime.gob.ve"
	curl := exec.Command("curl", url)

	output, err := curl.Output()
	if err != nil {
		fmt.Println(string(output))
	} else {
		fmt.Println("Hubo un error intetando hacer curl a " + url)
	}

	return err
}
