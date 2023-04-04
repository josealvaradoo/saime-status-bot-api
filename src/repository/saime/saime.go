package saime

import (
	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/storage/firestore/saime"
)

func Get() (model.Saime, error) {
	database := saime.Store{}
	return database.Get()
}

func Update(value string) (model.Saime, error) {
	database := saime.Store{}
	return database.Update(value)
}
