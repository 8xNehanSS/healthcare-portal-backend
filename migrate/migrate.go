package main

import (
	"healthcare-portal/initializers"
	"healthcare-portal/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// initializers.DB.AutoMigrate(&models.Login{})
	// initializers.DB.AutoMigrate(&models.User{})
	// initializers.DB.AutoMigrate(&models.News{})
	// initializers.DB.AutoMigrate(&models.Stories{})
	initializers.DB.AutoMigrate(&models.Appointment{})
}