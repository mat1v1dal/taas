package models

import "time"

type Exercise struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255"`
	Description string
	Category    string `gorm:"size:50"`
	CreatedAt   time.Time

	UserExercises []UserExercise
}
