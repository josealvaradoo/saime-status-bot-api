package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/josealvaradoo/saime-status-bot/src/model"
	"github.com/josealvaradoo/saime-status-bot/src/storage"
)

var ctx = context.Background()

func Get() (model.Saime, error) {
	value, err := storage.Cache().Get(ctx, "status").Result()

	if err != nil {
		return model.Saime{}, fiber.NewError(fiber.StatusNotFound, "error")
	}

	fmt.Println("Cache value:", value)

	return model.Saime{Status: value}, nil
}

func Update(value string) (model.Saime, error) {
	err := storage.Cache().Set(ctx, "status", value, 0).Err()

	if err != nil {
		return model.Saime{}, fiber.NewError(fiber.StatusBadGateway, "error")
	}

	return model.Saime{Status: value}, nil
}
