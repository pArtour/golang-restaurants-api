package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model

	Name    string `json:"name" gorm:"not null;text;unique"`
	Address string `json:"address" gorm:"not null;text"`
	Phone   string `json:"phone" gorm:"not null;text"`
	Email   string `json:"email" gorm:"not null;text"`
}
