package services

import (
	"errors"
	"fmt"
	"github.com/pArtour/golang-restaurants-api/pkg/database"
	"github.com/pArtour/golang-restaurants-api/pkg/models"
)

type RestaurantService struct {
	db *database.Database
}

func NewRestaurantService(db *database.Database) *RestaurantService {
	return &RestaurantService{db: db}
}

func (rs *RestaurantService) GetRestaurants() []models.Restaurant {
	restaurants := []models.Restaurant{}
	rs.db.Find(&restaurants)
	return restaurants
}

func (rs *RestaurantService) GetRestaurant(id int) (*models.Restaurant, error) {
	restaurant := new(models.Restaurant)
	err := rs.db.First(&restaurant, id)
	if err != nil {
		return nil, err.Error
	}
	return restaurant, nil
}

func (rs *RestaurantService) CreateRestaurant(restaurant *models.Restaurant) (*models.Restaurant, error) {
	r := new(models.Restaurant)
	err := rs.db.Where("name = ?", restaurant.Name).First(&r).Error
	if err == nil {
		return nil, errors.New("restaurant with this name already exists")
	}

	err = rs.db.Create(&restaurant).Error
	// print error
	fmt.Printf("error: %v", err)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}
