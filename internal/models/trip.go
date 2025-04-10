package models

import (
	"time"

	"gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	Longitude  float64
	Latitude   float64
	Province   string
	PostalCode int16
}

type Trip struct {
	gorm.Model
	DriverID       uint        `gorm:"not null"` // ID del conductor (UserDriver)
	Driver         *UserDriver `gorm:"foreignKey:DriverID"`
	OriginID       uint        `gorm:"not null"` // Lugar de origen
	Origin         *Place      `gorm:"foreignKey:OriginID"`
	DestinationID  uint        `gorm:"not null"` // Lugar de destino
	Destination    *Place      `gorm:"foreignKey:DestinationID"`
	DepartureTime  time.Time   `gorm:"not null"`
	AvailableSeats int         `gorm:"not null"`
	PricePerSeat   float64     `gorm:"not null"`
	Description    string
	Status         bool
	Cancelled      bool
	Participations []TripParticipation `gorm:"foreignKey:TripID"` // Relación con las participaciones
}

type TripParticipation struct {
	gorm.Model
	TripID      uint `gorm:"not null"`
	PassengerID uint `gorm:"not null"`
}
