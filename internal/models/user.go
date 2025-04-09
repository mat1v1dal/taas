package models

import (
	"gorm.io/gorm"
)

// Usuario base
type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Email            string `gorm:"unique;not null"`
	Password         string `gorm:"not null"`
	PassengerProfile *UserPassenger
	DriverProfile    *UserDriver
}

// Perfil de pasajero asociado a un usuario
type UserPassenger struct {
	gorm.Model
	UserID uint `gorm:"not null;uniqueIndex"` // Clave foránea única
	User   User `gorm:"foreignKey:UserID"`
	// Otros campos específicos para pasajeros
}

// Perfil de conductor asociado a un usuario
type UserDriver struct {
	gorm.Model
	UserID              uint     `gorm:"not null;uniqueIndex"` // Clave foránea única
	User                User     `gorm:"foreignKey:UserID"`
	DriverLicenseNumber string   `gorm:"not null"`
	Vehicle             *Vehicle // Relación con el vehículo
	// Otros campos específicos para conductores
}
