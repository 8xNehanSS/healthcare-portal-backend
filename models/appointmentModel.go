package models

import (
	"time"
)

type Appointment struct {
	ID        	uint
	PatientID 	uint `gorm:"not null"`
	DoctorID  	uint `gorm:"not null"`
	Date      	time.Time
	Reason    	string
	IsAccepted	bool `gorm:"default:false"`
	IsCompleted bool `gorm:"default:false"`
}
