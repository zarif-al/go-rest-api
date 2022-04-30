package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func EnvDbName() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_NAME")
}

func EnvDbUsername() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_USERNAME")
}

func EnvDbPassword() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_PASSWORD")
}

func EnvDbHost() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_HOST")
}

func EnvDbPort() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_PORT")
}

func ConnectDB() *gorm.DB {
	var err error

	var DB_USERNAME = EnvDbUsername()
	var DB_PASSWORD = EnvDbPassword()
	var DB_HOST = EnvDbHost()
	var DB_PORT = EnvDbPort()
	var DB_NAME = EnvDbName()

	// Set up Client
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// Client Instance
var DB *gorm.DB = ConnectDB()
