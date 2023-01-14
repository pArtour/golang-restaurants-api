package services

import "github.com/pArtour/golang-restaurants-api/pkg/database"

type Services struct {
	Restaurant *RestaurantService
}

func NewServices(db *database.Database) *Services {
	return &Services{
		Restaurant: NewRestaurantService(db),
	}
}
