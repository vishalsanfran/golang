package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video-api/controllers"
	"video-api/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "My Flix!"})
	})
	models.GetDbConnection()
	r.GET("/videos", controllers.GetAllVideos)
	r.POST("/videos", controllers.CreateVideo)
	r.Run()
}
