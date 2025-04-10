package services

import (
	"errors"
	"taas/internal/db"
	"taas/internal/models"
	"time"
)

type TripService struct{}

func (s *TripService) CreateTrip(trip *models.Trip) error {
	return db.DB.Create(trip).Error
}

func (s *TripService) GetTripByID(id uint) (*models.Trip, error) {
	var trip models.Trip
	err := db.DB.Preload("Origin").Preload("Destination").Preload("Driver").First(&trip, id).Error
	return &trip, err
}

func (s *TripService) GetActiveTripsByOriginAndDestination(originName string, destinationName string) ([]models.Trip, error) {
	var trips []models.Trip

	err := db.DB.
		Joins("JOIN places AS o ON o.id = trips.origin_id").
		Joins("JOIN places AS d ON d.id = trips.destination_id").
		Where("o.name = ? AND d.name = ?", originName, destinationName).
		Where("trips.status = ?", true). // si tenés campo status
		Preload("Origin").
		Preload("Destination").
		Preload("Driver").
		Find(&trips).Error

	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (s *TripService) VerifyTripStatus() error {
	var trips []models.Trip

	err := db.DB.Find(&trips).Error

	if err != nil {
		return err
	}

	now := time.Now()

	for _, trip := range trips {
		newStatus := true
		if trip.DepartureTime.After(now) {
			newStatus = false
			err := db.DB.Model(&trip).Update("status", newStatus).Error
			if err != nil {
				return err
			}
		}
		if trip.Cancelled && trip.Status {
			newStatus = false
			err := db.DB.Model(&trip).Update("status", newStatus).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *TripService) UpdateTripDate(Time time.Time, id string) error {
	if Time.After(time.Now()) {
		return nil
	}
	err := db.DB.Model(&models.Trip{}).Where("id = ?", id).Update("departure_time", Time).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *TripService) CancelTrip(id uint) error {
	return db.DB.Model(&models.Trip{}).Where("id = ?", id).Updates(map[string]interface{}{
		"cancelled": true,
		"status":    false,
	}).Error
}

func (s *TripService) GetTripsByDriver(driverName string, driverLastName string, driverDni string) ([]models.Trip, error) {
	var trips []models.Trip

	err := db.DB.
		Joins("JOIN user_drivers ON user_drivers.id = trips.driver_id").
		Joins("JOIN users ON users.id = user_drivers.user_id").
		Where("users.name = ? AND users.last_name = ? AND users.dni = ?", driverName, driverLastName, driverDni).
		Preload("Origin").
		Preload("Destination").
		Preload("Driver.User").
		Find(&trips).Error

	if err != nil {
		return nil, err
	}

	return trips, nil
}

func (s *TripService) JoinTrip(tripID uint, passengerID uint) error {
	var trip models.Trip

	// Verificar que el viaje existe y está activo
	err := db.DB.Preload("Participations").First(&trip, tripID).Error
	if err != nil {
		return errors.New("viaje no encontrado")
	}

	if !trip.Status || trip.Cancelled {
		return errors.New("el viaje no está activo o fue cancelado")
	}

	// Verificar si el pasajero ya está unido
	for _, participation := range trip.Participations {
		if participation.PassengerID == passengerID {
			return errors.New("el pasajero ya está unido a este viaje")
		}
	}

	// Verificar si hay asientos disponibles
	if trip.AvailableSeats <= 0 {
		return errors.New("no hay asientos disponibles")
	}

	// Agregar participación
	newParticipation := models.TripParticipation{
		TripID:      trip.ID,
		PassengerID: passengerID,
	}

	err = db.DB.Create(&newParticipation).Error
	if err != nil {
		return errors.New("no se pudo unir al viaje")
	}

	// Disminuir la cantidad de asientos disponibles
	trip.AvailableSeats -= 1
	err = db.DB.Save(&trip).Error
	if err != nil {
		return errors.New("error al actualizar asientos del viaje")
	}

	return nil
}
