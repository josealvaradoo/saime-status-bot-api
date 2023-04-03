package saime

import (
	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/storage/firestore"
)

type Store struct{}

func (s *Store) Get() (model.Saime, error) {
	snapshot, err := firestore.Client.Collection("saime").Doc("saime").Get(firestore.Ctx)

	if err != nil {
		return model.Saime{Status: ""}, model.ErrStatusNotExists
	}

	var m model.Saime
	snapshot.DataTo(&m)
	return model.Saime{Status: m.Status}, nil
}

func (s *Store) Update(value string) (model.Saime, error) {
	doc := firestore.Client.Collection("saime").Doc("saime")
	_, err := doc.Set(firestore.Ctx, model.Saime{Status: value})

	if err != nil {
		return model.Saime{}, model.ErrStatusCanNotBeUpdated
	}

	return model.Saime{Status: value}, nil
}
