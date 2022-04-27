package main

import (
	"web-services-gin/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	// Bind the router to each route from routes
	// services.AlbumServices(router)

	router.Run("localhost:8080")
}
