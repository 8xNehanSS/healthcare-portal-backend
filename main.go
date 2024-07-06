package main

import (
	"healthcare-portal/controllers"
	"healthcare-portal/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// entry points
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.CheckLogin)

	// users
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.DELETE("/users/:id", controllers.UserDelete)

	r.Run()
}