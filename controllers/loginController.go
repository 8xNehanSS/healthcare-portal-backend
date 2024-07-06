package controllers

import (
	"healthcare-portal/initializers"
	"healthcare-portal/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func CheckLogin(c *gin.Context) {
	var body struct {
		Username *string `json:"username"`
		Email *string `json:"email"`
		Password string `json:"password"`
		Type uint `json:"type"`
	}

	c.Bind(&body)
	if body.Username == nil && body.Email == nil {
		c.Status(400)
		return
	}

	var login models.Login
	if body.Username != nil {
		initializers.DB.Where("username = ?", *body.Username).First(&login)
	} else {
		initializers.DB.Where("email = ?", *body.Email).First(&login)
	}
	if login.ID == 0 {
		c.Status(400)
		return
	}
	if login.Password != body.Password {
		c.Status(400)
		return
	}
	if login.Type != body.Type { 
		c.Status(400)
		return
	}
	c.Status(200)
}

func CreateLogin(username string, email string, password string, userType uint, userID uint, user models.User) error {
	login := models.Login{
		UserID:   userID,
		User:     user,
		Username: username,
		Email:    email,
		Password: password,
		Type:     userType,
	}
	
	initializers.DB.Create(&login)
	return nil
}