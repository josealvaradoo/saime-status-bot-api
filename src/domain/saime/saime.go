package saime

import (
	"fmt"
	"net/http"

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
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Hubo un error intetando hacer get a " + url)
	}

	return err
}
