package controllers

import (
	"healthcare-portal/initializers"
	"healthcare-portal/models"

	"github.com/gin-gonic/gin"
)

func ValidateUsername (c *gin.Context) {
	var body struct {
		Username string `json:"username"`
	}
	c.Bind(&body)
	var login models.Login
	if err := initializers.DB.Where("username = ?", body.Username).First(&login).Error; err != nil {
		c.Status(400)
		return
	}
	c.Status(200)
}

// func CreatePatientUser (c *gin.Context) {
// 	var body struct {
// 		FirstName string `json:"firstName"`
// 		LastName string `json:"lastName"`
// 		Address string `json:"address"`
// 		Email string `json:"email"`
// 		Age uint8 `json:"age"`
// 		Birthday *time.Time `json:"birthday"`
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 		Type uint `json:"type"`
// 	}

// 	c.Bind(&body)

// 	patient := models.Patient{
// 		FirstName: body.FirstName,
// 		LastName: body.LastName,
// 		Address: body.Address,
// 		Email: body.Email,
// 		Age: body.Age,
// 		Birthday: body.Birthday,
// 	}
// 	result := initializers.DB.Create(&patient)
// 	if result.Error != nil {
// 		c.Status(400)
// 		return
// 	}
// 	err := CreateLogin(body.Username, body.Email, body.Password, body.Type, patient.ID, &patient, nil)
// 	if err != nil {
// 		c.Status(400)
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"user": patient,
// 	})
// }

// func GetUsers (c *gin.Context) {
// 	var users []models.User
// 	initializers.DB.Find(&users)
// 	c.JSON(200, gin.H{
// 		"users": users,
// 	})
// }

// func GetUser (c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
// 	initializers.DB.First(&user, id)
// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }

// func UserUpdate (c *gin.Context) {
// 	id := c.Param("id")
// 	var body struct {
// 		FirstName string `json:"first_name"`
// 		LastName string `json:"last_name"`
// 		Address string `json:"address"`
// 		Email string `json:"email"`
// 		Age uint8 `json:"age"`
// 		Birthday *time.Time `json:"birthday"`
// 	}

// 	c.Bind(&body)

// 	var user models.User
// 	initializers.DB.First(&user, id)

// 	userUpdated := models.User{
// 		FirstName: body.FirstName,
// 		LastName: body.LastName,
// 		Address: body.Address,
// 		Email: body.Email,
// 		Age: body.Age,
// 		Birthday: body.Birthday,
// 	}

// 	initializers.DB.Model(&user).Updates(userUpdated)

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }

// func UserDelete (c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
// 	initializers.DB.First(&user, id)
// 	initializers.DB.Delete(&user)
// 	c.Status(200)
// }