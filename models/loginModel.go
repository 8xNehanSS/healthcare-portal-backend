package models

// type = 9 -> admin
// type = 1 -> doctor
// type = 2 -> patient

type Login struct {
	ID       uint
	UserID   uint `gorm:"not null"`
	User     User `gorm:"foreignKey:UserID;references:ID"`
	Username string
	Email    string
	Password string
	Type     uint
}