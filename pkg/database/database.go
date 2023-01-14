package database

import (
	"fmt"
	"github.com/pArtour/golang-restaurants-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database connected")

	log.Print("Running migrations")
	db.AutoMigrate(&models.Restaurant{})
	return &Database{db}
}
