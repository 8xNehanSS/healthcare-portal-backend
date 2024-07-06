package main

import (
	"healthcare-portal/controllers"
	"healthcare-portal/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/register", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.Run()
}