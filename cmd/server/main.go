package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/golang-restaurants-api/pkg/database"
	"github.com/pArtour/golang-restaurants-api/pkg/handlers"
	"github.com/pArtour/golang-restaurants-api/pkg/services"
)

func main() {
	app := fiber.New()

	db := database.NewDatabase()
	srv := services.NewServices(db)

	handlers.NewHandlers(srv, app)

	app.Listen(":3000")
}
