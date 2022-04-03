package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnvMongoDBName() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_DB_NAME")
}

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}

func ConnectDB() *mongo.Client {
	// Set up Client
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	// Set up context with timeout for connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Attempt Connection
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Ping DB
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

// Client Instance
var DB *mongo.Client = ConnectDB()

//Getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var MONGO_DB_NAME = EnvMongoDBName()
	collection := client.Database(MONGO_DB_NAME).Collection(collectionName)
	return collection
}
