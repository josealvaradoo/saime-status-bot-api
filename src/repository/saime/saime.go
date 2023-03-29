package saime

import (
	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/storage/memory/saime"
)

func Get() (model.Saime, error) {
	cache := saime.Store{}
	return cache.Get()
}

func Update(value string) (model.Saime, error) {
	cache := saime.Store{}
	return cache.Update(value)
}
