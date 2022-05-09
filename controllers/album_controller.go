package controllers

import (
	"log"
	"net/http"
	"web-services-gin/configs"
	"web-services-gin/dtos"
	"web-services-gin/models"

	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

func GetAlbum(c *gin.Context) {
	var album models.Album
	id, _ := c.Params.Get("id")
	if err := configs.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, album)
		return
	}
}

func CreateAlbum(c *gin.Context) {
	var album models.Album
	// Bind input in context body to album pointer
	if err := c.BindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := gonanoid.New()
	if err != nil {
		log.Println(color.Ize(color.Red, "Error : "+err.Error()))
	}

	album.ID = id

	if err := configs.DB.Create(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, album)
		return
	}
}

func InsertAlbums(albums []models.Album) dtos.ErrorDTO {
	var newError dtos.ErrorDTO
	if err := configs.DB.Create(&albums).Error; err != nil {
		newError.Error = true
		newError.Message = err.Error()
	} else {
		newError.Error = false
		newError.Message = "Album inserted successfully"
	}
	return newError

}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	id, _ := c.Params.Get("id")
	if err := configs.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		if err := c.BindJSON(&album); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		configs.DB.Save(&album)
		c.JSON(200, album)
		return
	}
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album
	id, _ := c.Params.Get("id")
	if err := configs.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		configs.DB.Delete(&album)
		c.JSON(200, album)
		return
	}
}
