package main

import (
	"web-services-gin/configs"
	"web-services-gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	// Bind the router to each route from routes
	services.AlbumsServices(router)
	services.UploadServices(router)

	router.Run("localhost:8080")
}
