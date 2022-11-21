package services

import (
	"web-services-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumsServices(router *gin.Engine) {
	router.GET("/get-all-albums", controllers.GetAlbums)
	router.POST("/create-album", controllers.CreateAlbum)
	router.GET("/get-album/:id", controllers.GetAlbum)
	router.PATCH("/update-album/:id", controllers.UpdateAlbum)
	router.DELETE("/delete-album/:id", controllers.DeleteAlbum)
}
