package saime

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/storage/cache"
)

var ctx = context.Background()

func Get() (model.Saime, error) {
	value, err := cache.Cache().Get(ctx, "status").Result()

	if err != nil {
		return model.Saime{}, fiber.NewError(fiber.StatusNotFound, "error")
	}

	return model.Saime{Status: value}, nil
}

func Update(value string) (model.Saime, error) {
	err := cache.Cache().Set(ctx, "status", value, 0).Err()

	if err != nil {
		return model.Saime{}, fiber.NewError(fiber.StatusBadGateway, "error")
	}

	return model.Saime{Status: value}, nil
}
