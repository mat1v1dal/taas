package models

import "time"

type Emotion struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"not null"`
	Mood           string `gorm:"type:enum('happy','sad','angry','anxious','neutral');not null"`
	Note           string
	Intensity      int    `gorm:"type:enum('1','2','3','4','5','6','7','8','9','10');not null"`
	ContextualInfo string `gorm:"type:text"`
	CreatedAt      time.Time
}
