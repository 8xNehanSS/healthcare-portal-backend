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

func GetPublicData(c *gin.Context) {
	var news []models.News
	initializers.DB.Find(&news)
	var stories []models.Stories
	initializers.DB.Find(&stories)
	c.JSON(200, gin.H{
		"news":    news,
		"stories": stories,
	})
}