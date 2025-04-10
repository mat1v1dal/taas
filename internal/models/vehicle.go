package models

import (
	"database/sql/driver"
	"errors"

	"gorm.io/gorm"
)

// models/vehicle.go
type Vehicle struct {
	gorm.Model
	Make   string `gorm:"not null"`
	Models string `gorm:"not null"` // Cambiado de int a string
	Plate  string `gorm:"unique;not null"`
}

func (v *Vehicle) Value() (driver.Value, error) {
	if v == nil {
		return nil, nil
	}
	return v.ID, nil
}

// Escanear desde la base de datos
func (v *Vehicle) Scan(value interface{}) error {
	if value == nil {
		v = nil
		return nil
	}
	if id, ok := value.(int64); ok {
		v.ID = uint(id)
		return nil
	}
	return errors.New("failed to scan Vehicle")
}
