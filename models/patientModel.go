package models

import "time"

// Patient struct represents a patient in the system.
type Patient struct {
	PatientID     string    `gorm:"primaryKey;type:varchar(255);not null"` // Unique identifier for the patient
	FirstName     string    `gorm:"size:100;not null"`                       // First name of the patient
	LastName      string    `gorm:"size:100;not null"`                       // Last name of the patient
	DateOfBirth   time.Time `gorm:"not null"`                                // Date of birth of the patient
	Gender        string    `gorm:"size:10"`                                 // Gender (e.g., Male, Female, Other)
	ContactNumber string    `gorm:"size:15"`                                 // Contact number as a string
	Address       string    `gorm:"size:255"`                                // Address of the patient
	CreatedAt     time.Time `gorm:"autoCreateTime"`                          // Automatically set the created timestamp
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`                          // Automatically set the updated timestamp
}
