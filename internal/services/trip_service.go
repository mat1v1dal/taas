package services

import (
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
			err := db.DB.Model(&trip).Update("status", newStatus)
			if err != nil {
				return err.Error
			}
		}
		if trip.Cancelled && trip.Status {
			newStatus = false
			err := db.DB.Model(&trip).Update("status", newStatus)
			if err != nil {
				return err.Error
			}
		}
	}
	return nil
}

func (s *TripService) UpdateTripDate(time time.Time, id s)
