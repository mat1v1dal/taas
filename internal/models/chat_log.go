package models

import "time"

type ChatLog struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Sender    string `gorm:"type:enum('user','ai');not null"`
	Message   string `gorm:"type:text;not null"`
	CreatedAt time.Time
}
