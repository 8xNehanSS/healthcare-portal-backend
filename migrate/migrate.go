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
	initializers.DB.AutoMigrate(&models.Login{})
	initializers.DB.AutoMigrate(&models.User{})
}