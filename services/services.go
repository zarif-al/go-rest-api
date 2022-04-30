package services

import (
	"web-services-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumsServices(router *gin.Engine) {
	//All services related to albums comes here
	// router.GET("/get-album/:albumId", controllers.GetAlbum())
	router.GET("/get-all-albums", controllers.GetAlbums)
}
