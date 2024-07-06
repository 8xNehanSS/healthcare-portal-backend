package controllers

import (
	"healthcare-portal/initializers"
	"healthcare-portal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": login.ID,
	"exp": time.Now().Add(time.Hour * 24 *5).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.Status(500)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*5, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
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

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}