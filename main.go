package main

import (
	"web-services-gin/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	// Bind the router to each route from routes

	router.Run("localhost:8080")
}
