package services

import (
	"web-services-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumServices(router *gin.Engine) {
	//All routes related to albums comes here
	router.GET("/get-album/:albumId", controllers.GetAlbum())
	router.POST("/url-status", controllers.CheckStatus())
	router.GET("/get-all-albums", controllers.GetAllAlbums())
	router.POST("/create-album", controllers.CreateAlbum())
	router.PUT("/edit-album/:albumId", controllers.EditAlbum())
	router.DELETE("/delete-album/:albumId", controllers.DeleteAlbum())
}
