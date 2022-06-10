package main

import (
	"log"
	"os"
	"web-services-gin/configs"
	"web-services-gin/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()

	configs.ConnectDB()
	router.Use(cors.Default())
	//routes
	services.AlbumServices(router)

	log.Printf("Running on port %s", port)

	router.Run("localhost:" + port)
}
