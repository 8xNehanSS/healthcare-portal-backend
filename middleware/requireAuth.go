package middleware

import (
	"fmt"
	"healthcare-portal/initializers"
	"healthcare-portal/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Cookie not found")
		c.AbortWithStatus(401)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(401)

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(401)
		}
		var user models.Login
		if err := initializers.DB.Where("user_id = ?", claims["sub"]).First(&user).Error; err != nil {
			c.Status(401)
			return
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}