package configs

import (
	"log"
	"os"

	"web-services-gin/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func EnvDbName() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: DATABASE_NAME : %v", err)
	}

	return os.Getenv("DATABASE_NAME")
}

func EnvDbUsername() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: DATABASE_USERNAME")
	}

	return os.Getenv("DATABASE_USERNAME")
}

func EnvDbPassword() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: DATABASE_PASSWORD")
	}

	return os.Getenv("DATABASE_PASSWORD")
}

func EnvDbHost() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: DATABASE_HOST")
	}

	return os.Getenv("DATABASE_HOST")
}

func EnvDbPort() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: DATABASE_PORT")
	}

	return os.Getenv("DATABASE_PORT")
}

func ConnectDB() *gorm.DB {
	var err error

	var DB_NAME = EnvDbName()
	var DB_PASSWORD = EnvDbPassword()
	var DB_USERNAME = EnvDbUsername()
	var DB_HOST = EnvDbHost()
	var DB_PORT = EnvDbPort()

	// Set up Client
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 500, Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Hello")
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Album{})

	return db
}

// Client Instance
var DB *gorm.DB = ConnectDB()
