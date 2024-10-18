package main

import (
	"healthcare-portal/controllers"
	"healthcare-portal/initializers"
	"healthcare-portal/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3001"}, // Allow frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true, // Allow cookies
        MaxAge:           12 * time.Hour,
    }))

	// entry points
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	// r.POST("/register",middleware.RequireAuth, controllers.CreatePatientUser)
	r.POST("/login", controllers.CheckLogin)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/validate-username", middleware.RequireAuth, controllers.ValidateUsername)
	r.GET("/getpublic", controllers.GetPublicData);

	// users
	// r.GET("/users",middleware.RequireAuth, controllers.GetUsers)
	// r.GET("/users/:id",middleware.RequireAuth, controllers.GetUser)
	// r.PUT("/users/:id",middleware.RequireAuth, controllers.UserUpdate)
	// r.DELETE("/users/:id",middleware.RequireAuth, controllers.UserDelete) //check

	// doctors
	r.GET("/doctor/getdashboard", middleware.RequireAuth, controllers.GetDocDashboardData)

	// patients
	// r.POST("/patient/createappointment", middleware.RequireAuth, controllers.CreateAppointment)

	r.Run()
}