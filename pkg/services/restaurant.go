package services

import (
	"github.com/pArtour/golang-restaurants-api/pkg/database"
	"github.com/pArtour/golang-restaurants-api/pkg/models"
)

type RestaurantService struct {
	db *database.Database
}

func NewRestaurantService(db *database.Database) *RestaurantService {
	return &RestaurantService{db: db}
}

func (r *RestaurantService) GetRestaurants() []models.Restaurant {
	return []models.Restaurant{}
}
