package models

import "time"

// Login struct represents user login information in the system.
type Login struct {
	UserID    string    `gorm:"primaryKey;type:varchar(255);not null"` // Unique identifier for the user
	Username  string    `gorm:"size:100;unique;not null"`               // Unique username for the user
	Email     string    `gorm:"size:100;unique;not null"`               // Unique email for the user
	Password  string    `gorm:"not null"`                                // Password for security
	Type      uint      `gorm:"not null"`                                // User type: 1 = doctor, 2 = patient, 9 = admin
	CreatedAt time.Time `gorm:"autoCreateTime"`                          // Automatically set the created timestamp
	UpdatedAt time.Time `gorm:"autoUpdateTime"`                          // Automatically set the updated timestamp
}
