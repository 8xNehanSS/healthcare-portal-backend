package main

import (
	"healthcare-portal/controllers"
	"healthcare-portal/initializers"
	"healthcare-portal/middleware"

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
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// users
	r.GET("/users",middleware.RequireAuth, controllers.GetUsers)
	r.GET("/users/:id",middleware.RequireAuth, controllers.GetUser)
	r.PUT("/users/:id",middleware.RequireAuth, controllers.UserUpdate)
	r.DELETE("/users/:id",middleware.RequireAuth, controllers.UserDelete)

	r.Run()
}