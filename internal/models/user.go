package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100"`
	Email        string `gorm:"size:100;unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"type:enum('user','therapist');default:'user'"`
	CreatedAt    time.Time

	Emotions       []Emotion
	ChatLogs       []ChatLog
	Sessions       []Session `gorm:"foreignKey:UserID"`
	TherapistSlots []Session `gorm:"foreignKey:TherapistID"`
	UserExercises  []UserExercise
}
