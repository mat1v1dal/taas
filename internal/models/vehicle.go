package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model `swaggerignore:"true"`
	Make       string `gorm:"not null"`
	Models     string `gorm:"not null"`
	Plate      string `gorm:"unique;not null"`
}
