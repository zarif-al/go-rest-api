package controllers

import (
	"net/http"
	"web-services-gin/configs"
	"web-services-gin/models"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	if err := configs.DB.Find(&albums).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, albums)
		return
	}
}

func CreateAlbum(c *gin.Context) {
	var album models.Album
	if err := c.BindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := configs.DB.Create(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, album)
		return
	}
}
