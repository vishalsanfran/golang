package controllers

import (
	"net/http"
	"video-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllVideos(c *gin.Context) {
	var videos []models.Video
	models.DB.Find(&videos)
	c.JSON(http.StatusOK, gin.H{"data": videos})
}

func CreateVideo(c *gin.Context) {
	var input models.VideoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	video := models.Video{Title: input.Title, Duration: input.Duration, Category: input.Category}
	models.DB.Create(&video)
	c.JSON(http.StatusOK, gin.H{"data": video})
}
