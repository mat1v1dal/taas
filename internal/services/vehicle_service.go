package services

import (
	"errors"
	"taas/internal/db"
	"taas/internal/models"
)

type VehicleService struct{}

func (s *VehicleService) CreateVehicle(vehicle *models.Vehicle) error {
	err := db.DB.Create(vehicle).Error
	if err != nil {
		return errors.New("error creating vehicle")
	}
	return nil
}

func (s *VehicleService) GetVehicleByID(id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := db.DB.First(&vehicle, id).Error
	if err != nil {
		return nil, errors.New("vehicle not found")
	}
	return &vehicle, nil
}
