package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Make   string `gorm:"not null"`
	Modelo int
	Plate  string `gorm:"unique;not null"`
}
