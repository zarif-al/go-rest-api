package controllers

import (
	"context"
	"net/http"
	"time"
	"web-services-gin/configs"
	"web-services-gin/models"
	dtos "web-services-gin/dtos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var albumCollection *mongo.Collection = configs.GetCollection(configs.DB, "albums")

var validate = validator.New()

func CreateAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var album models.Album
		defer cancel()

		// bind input from context to album and validate the request body
		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, dtos.ResponseDTO{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// use validator library to validate required fields
		if validationErr := validate.Struct(&album); validationErr != nil {
			c.JSON(http.StatusBadRequest, dtos.ResponseDTO{Status: http.StatusBadRequest, Message: "Invalid request body", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAlbum := models.Album{
			ID:     primitive.NewObjectID(),
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		}

		result, err := albumCollection.InsertOne(ctx, newAlbum)

		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{Status: http.StatusInternalServerError, Message: "Error inserting album", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, dtos.ResponseDTO{Status: http.StatusCreated, Message: "Album created successfully", Data: map[string]interface{}{"data": result}})

	}
}

func GetAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		albumId := c.Param("albumId")

		var album models.Album
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(albumId)

		// Attempt to fetch album by id, if successfull bind to album var
		err := albumCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&album)

		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Error fetching album",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Return album
		c.JSON(http.StatusOK, dtos.ResponseDTO{
			Status:  http.StatusOK,
			Message: "Album fetched successfully",
			Data:    map[string]interface{}{"data": album},
		})
	}
}
