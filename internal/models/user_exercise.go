package models

import "time"

type UserExercise struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	ExerciseID  uint   `gorm:"not null"`
	Status      string `gorm:"type:enum('pending','completed');default:'pending'"`
	CompletedAt *time.Time
}
