package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/golang-restaurants-api/pkg/services"
)

type Restaurant struct {
	service *services.RestaurantService
	app     *fiber.App
}

func NewRestaurantHandler(service *services.RestaurantService, app *fiber.App) *Restaurant {
	r := &Restaurant{service: service, app: app}
	r.RegisterRoutes()
	return r
}

func (r *Restaurant) RegisterRoutes() {
	r.app.Get("/restaurants", r.GetRestaurants)
}

func (r *Restaurant) GetRestaurants(c *fiber.Ctx) error {
	return r.service.GetRestaurants()
}
