package saime

import (
	"github.com/josealvaradoo/saime-status-bot/src/model"
	memory "github.com/josealvaradoo/saime-status-bot/src/storage/memory"
)

var mem = memory.NewMemory()

type Store struct{}

func (s Store) Get() (model.Saime, error) {
	saime, err := mem.Get()

	if err != nil {
		return model.Saime{Status: ""}, err
	}

	return saime, nil
}

func (s Store) Update(value string) (model.Saime, error) {
	err := mem.Post(value)

	if err != nil {
		return model.Saime{}, err
	}

	return model.Saime{Status: value}, nil
}
