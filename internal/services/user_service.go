package services

import (
	"errors"
	"taas/internal/db"
	"taas/internal/models"
)

type UserService struct{}

func (s *UserService) CreateUser(user *models.User) error {

	// Verificar si el usuario ya existe
	var existingUser models.User
	err := db.DB.Where("email = ? OR dni = ?", user.Email, user.Dni).First(&existingUser).Error
	if err == nil {
		return errors.New("user already exists")
	}

	err = db.DB.Create(user).Error
	if err != nil {
		return errors.New("error creating user")
	}
	return nil
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

func (s *UserService) RegisterDriver(user *models.User, driverLicense string) error {
	// Crear el usuario base
	err := s.CreateUser(user)
	if err != nil {
		return err
	}

	driver := &models.UserDriver{
		UserID:              user.ID,
		DriverLicenseNumber: driverLicense,
	}

	return db.DB.Create(driver).Error
}

func (s *UserService) RegisterVehicleToDriver(driverID uint, vehicleID uint) error {
	var driver models.UserDriver
	vs := &VehicleService{}
	vehicle, err := vs.GetVehicleByID(vehicleID)

	if err != nil {
		return errors.New("vehicle not found")
	}

	err = db.DB.First(&driver, driverID).Error

	if err != nil {
		return errors.New("driver not found")
	}

	driver.VehicleID = vehicleID
	driver.Vehicle = vehicle
	return db.DB.Save(&driver).Error
}
