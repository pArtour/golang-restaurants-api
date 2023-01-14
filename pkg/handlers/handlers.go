package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/golang-restaurants-api/pkg/services"
)

type Handlers struct {
	RestaurantHandler *RestaurantHandler
}

func NewHandlers(service *services.Services, app *fiber.App) *Handlers {
	return &Handlers{
		RestaurantHandler: NewRestaurantHandler(service.Restaurant, app),
	}
}
