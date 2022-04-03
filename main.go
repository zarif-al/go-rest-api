package main

import (
	"web-services-gin/configs"
	"web-services-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	//routes
	routes.AlbumsRoute(router)

	router.Run("localhost:8080")
}
