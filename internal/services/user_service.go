package services

import (
	"errors"
	"taas/internal/db"
	"taas/internal/models"
)

type UserService struct{}

func (s *UserService) CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (s *UserService) RegisterPassenger(user *models.User) error {
	// Crear el usuario base
	err := s.CreateUser(user)
	if err != nil {
		return err
	}

	// Crear perfil de pasajero
	passenger := &models.UserPassenger{
		UserID: user.ID,
	}
	return db.DB.Create(passenger).Error
}

func (s *UserService) RegisterDriver(user *models.User, driverLicense string, vehicle *models.Vehicle) error {
	// Crear el usuario base
	err := s.CreateUser(user)
	if err != nil {
		return err
	}

	// Crear vehículo
	if vehicle != nil {
		err = db.DB.Create(vehicle).Error
		if err != nil {
			return errors.New("no se pudo guardar el vehículo")
		}
	}

	// Crear perfil de conductor
	driver := &models.UserDriver{
		UserID:              user.ID,
		DriverLicenseNumber: driverLicense,
		Vehicle:             vehicle,
	}
	return db.DB.Create(driver).Error
}
