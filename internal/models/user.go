package models

import (
	"gorm.io/gorm"
)

// Usuario base
type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	LastName         string `gorm:"not null"`
	Email            string `gorm:"unique;not null"`
	Password         string `gorm:"not null"`
	Phone            string `gorm:"not null"`
	Dni              string `gorm:"not null"`
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
	UserID              uint     `gorm:"not null;uniqueIndex"`
	User                User     `gorm:"foreignKey:UserID"`
	DriverLicenseNumber string   `gorm:"not null"`
	VehicleID           *uint    // Cambiado a puntero para permitir null
	Vehicle             *Vehicle `gorm:"foreignKey:VehicleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
