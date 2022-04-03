package routes

import (
	"web-services-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumsRoute(router *gin.Engine) {
	//All routes related to albums comes here
	router.POST("/create-album", controllers.CreateAlbum())
	router.GET("/album/:albumId", controllers.GetAlbum())
}
