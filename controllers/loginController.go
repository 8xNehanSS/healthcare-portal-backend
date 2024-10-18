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
	if err := initializers.DB.Where("username = ?", *body.Username).First(&login).Error; err != nil {
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
	"sub": login.UserID,
	"exp": time.Now().Add(time.Hour * 24 *5).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.Status(500)
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 3600*24*5, "/", "", true, false)
	c.JSON(http.StatusOK, gin.H{})
}

// func CreateLogin(username string, email string, password string, userType uint, userID string, userPatient *models.Patient, userDoctor *models.Doctor) error {
// 	if(userPatient != nil && userDoctor != nil) {
// 		return nil
// 	}
// 	if(userType == 1) {
// 		login := models.Login{
// 			UserID:   userID,
// 			Username: username,
// 			Email:    email,
// 			Password: password,
// 			Type:     userType,
// 			Patient: userPatient,
// 			PatientID: &userID,
// 		}
		
// 		initializers.DB.Create(&login)
// 		return nil
// 	} else {
// 		login := models.Login{
// 			UserID:   userID,
// 			Username: username,
// 			Email:    email,
// 			Password: password,
// 			Type:     userType,
// 			Doctor: userDoctor,
// 			DoctorID: &userID,
// 		}
		
// 		initializers.DB.Create(&login)
// 		return nil
// 	}
// }

func Validate(c *gin.Context) {
	var firstName, lastName string
	login, _ := c.Get("user")
	loginData, _ := login.(models.Login)
	if(loginData.Type == 1) {
		var user models.Doctor
		if err := initializers.DB.Where("doctor_id = ?", loginData.UserID).First(&user).Error; err != nil {
			c.Status(401)
			return
		}
		firstName = user.FirstName
		lastName = user.LastName
	} else {
		var user models.Patient
		if err := initializers.DB.Where("patient_id = ?", loginData.UserID).First(&user).Error; err != nil {
			c.Status(401)
			return
		}
		firstName = user.FirstName
		lastName = user.LastName
	}

	c.JSON(http.StatusOK, gin.H{
		"userID": loginData.UserID,
		"username": loginData.Username,
		"firstname": firstName,
		"lastname": lastName,
		"loginType": loginData.Type,
	})
}