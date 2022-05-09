package main

import (
	"log"
	"os"
	"web-services-gin/configs"
	"web-services-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func EnvGetPort() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: PORT")
	}

	return os.Getenv("PORT")
}

func main() {
	router := gin.Default()

	configs.ConnectDB()

	// Bind the router to each route from routes
	services.AlbumsServices(router)
	services.UploadServices(router)

	router.Run(":" + EnvGetPort())
}
