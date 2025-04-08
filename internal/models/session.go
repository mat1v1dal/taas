package models

import "time"

type Session struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	TherapistID uint      `gorm:"not null"`
	ScheduledAt time.Time `gorm:"not null"`
	Link        string
	Notes       string
	CreatedAt   time.Time
}
