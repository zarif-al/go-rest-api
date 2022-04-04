package main

import (
	"web-services-gin/configs"
	"web-services-gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	//routes
	services.AlbumServices(router)

	router.Run("localhost:8080")
}
