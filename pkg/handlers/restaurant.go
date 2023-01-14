package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/golang-restaurants-api/pkg/models"
	"github.com/pArtour/golang-restaurants-api/pkg/services"
	"gorm.io/gorm"
	"strconv"
)

// RestaurantHandler is a handler for restaurant routes
type RestaurantHandler struct {
	service *services.RestaurantService
	app     *fiber.App
}

// NewRestaurantHandler creates a new restaurant handler
func NewRestaurantHandler(service *services.RestaurantService, app *fiber.App) *RestaurantHandler {
	rh := &RestaurantHandler{service: service, app: app}
	rh.RegisterRoutes()
	return rh
}

// RegisterRoutes registers all routes for the restaurant handler
func (rh *RestaurantHandler) RegisterRoutes() {
	rh.app.Get("/restaurants", rh.GetRestaurants)
	rh.app.Post("/restaurants", rh.CreateRestaurant)
	rh.app.Get("/restaurants/:id", rh.GetRestaurant)
}

// GetRestaurants returns all restaurants
func (rh *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	restaurants := rh.service.GetRestaurants()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"restaurants": restaurants,
		"error":       nil,
	})
}

// GetRestaurant returns a restaurant by id
func (rh *RestaurantHandler) GetRestaurant(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "invalid id",
			"restaurant": nil,
		})
	}
	restaurant, err := rh.service.GetRestaurant(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      err.Error(),
			"restaurant": nil,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      err.Error(),
			"restaurant": nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":      nil,
		"restaurant": restaurant,
	})
}

// CreateRestaurant creates a new restaurant
func (rh *RestaurantHandler) CreateRestaurant(c *fiber.Ctx) error {
	r := new(models.Restaurant)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      err.Error(),
			"restaurant": nil,
		})
	}
	restaurant, err := rh.service.CreateRestaurant(r)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      err.Error(),
			"restaurant": nil,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":      nil,
		"restaurant": restaurant,
	})
}
