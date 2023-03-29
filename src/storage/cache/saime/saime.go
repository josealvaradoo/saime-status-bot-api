package saime

import (
	"context"

	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/storage/cache"
)

var ctx = context.Background()

type Store struct{}

func (s Store) Get() (model.Saime, error) {
	status, err := cache.Cache().Get(ctx, "status").Result()
	if err != nil {
		return model.Saime{Status: ""}, model.ErrStatusNotExists
	}

	return model.Saime{Status: status}, nil
}

func (s Store) Update(value string) (model.Saime, error) {
	err := cache.Cache().Set(ctx, "status", value, 0).Err()

	if err != nil {
		return model.Saime{}, model.ErrStatusCanNotBeUpdated
	}

	return model.Saime{Status: value}, nil
}
