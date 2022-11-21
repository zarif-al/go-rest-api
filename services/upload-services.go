package services

import (
	"web-services-gin/controllers"

	"github.com/gin-gonic/gin"
)

func UploadServices(router *gin.Engine) {
	router.POST("/upload-file", controllers.UploadFile)
}
