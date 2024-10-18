package models

import (
	"time"
)

type Appointment struct {
	ID        string             `gorm:"primaryKey"` // Changed to string for alphanumeric IDs
	PatientID string             `gorm:"not null"`
	DoctorID  string             `gorm:"not null"`
	Date      time.Time          `gorm:"not null"`
	Reason    string             `gorm:"size:255"` // Adjust size as needed
	Status    string   			 `gorm:"not null"` // Set the type according to your database
	CreatedAt time.Time          `gorm:"autoCreateTime"`
	UpdatedAt time.Time          `gorm:"autoUpdateTime"`
}
