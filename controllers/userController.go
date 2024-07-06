package controllers

import (
	"healthcare-portal/initializers"
	"healthcare-portal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser (c *gin.Context) {
	var body struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Address string `json:"address"`
		Email *string `json:"email"`
		Age uint8 `json:"age"`
		Birthday *time.Time `json:"birthday"`
	}

	c.Bind(&body)

	user := models.User{
		FirstName: body.FirstName,
		LastName: body.LastName,
		Address: body.Address,
		Email: body.Email,
		Age: body.Age,
		Birthday: body.Birthday,
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers (c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}