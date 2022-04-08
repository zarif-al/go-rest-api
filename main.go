package main

import (
	"log"
	"os"
	"web-services-gin/configs"
	"web-services-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getPortNumber() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("PORT")
}

func main() {
	router := gin.Default()

	configs.ConnectDB()

	//routes
	services.AlbumServices(router)

	var port = getPortNumber()

	router.Run(":" + port)
}
