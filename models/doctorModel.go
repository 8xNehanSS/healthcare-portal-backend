package models

import (
	"time"
)

// Doctor struct represents a medical doctor in the system.
type Doctor struct {
	DoctorID     string    `gorm:"primaryKey;type:varchar(255);not null"` // Unique identifier for the doctor
	FirstName    string    `gorm:"size:100;not null"`                       // First name of the doctor
	LastName     string    `gorm:"size:100;not null"`                       // Last name of the doctor
	Specialty    string    `gorm:"size:100"`                                // Medical specialty of the doctor
	ContactNumber string    `gorm:"size:15"`                                 // Contact number as a string for better formatting
	Address       string    `gorm:"size:255"`                                // Address of the doctor
	Bio           string    `gorm:"size:500"`                                // Brief biography of the doctor
	CreatedAt     time.Time `gorm:"autoCreateTime"`                          // Automatically set the created timestamp
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`                          // Automatically set the updated timestamp
}
