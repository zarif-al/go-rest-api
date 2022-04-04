package controllers

import (
	"context"
	"net/http"
	"time"
	"web-services-gin/configs"
	dtos "web-services-gin/dtos"
	"web-services-gin/models"

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

func EditAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		albumId := c.Param("albumId")
		var album dtos.UpdateAlbum
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(albumId)

		// Validate Request Body
		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, dtos.ResponseDTO{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body, bind failed",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// use validator library to validate required fields
		if validationErr := validate.Struct(&album); validationErr != nil {
			c.JSON(http.StatusBadRequest, dtos.ResponseDTO{Status: http.StatusBadRequest,
				Message: "Invalid request body, validation failed",
				Data:    map[string]interface{}{"data": validationErr.Error()},
			})
			return
		}

		// Check if album exists

		var existingAlbum models.Album
		findErr := albumCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&existingAlbum)

		if findErr != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.ResponseDTO{
					Status:  http.StatusInternalServerError,
					Message: "Album not found!",
					Data:    map[string]interface{}{"data": findErr.Error()},
				},
			)
			return
		}

		// Attempt to update album by id, if successfull bind to album var
		update := bson.M{"$set": album}

		result, err := albumCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Error updating album",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Return updated album
		var updatedAlbum models.Album

		if result.MatchedCount == 1 {
			err := albumCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAlbum)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{Status: http.StatusInternalServerError,
					Message: "Error Fetching Updated Album",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
		}

		c.JSON(http.StatusOK, dtos.ResponseDTO{Status: http.StatusOK,
			Message: "Album updated successfully",
			Data:    map[string]interface{}{"data": updatedAlbum},
		})
	}
}

func DeleteAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		albumId := c.Param("albumId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(albumId)

		// Attempt to delete album by id, if successfull bind to album var
		result, err := albumCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Error deleting album",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Check if album deleted successfully and return
		if result.DeletedCount < 1 {
			c.JSON(
				http.StatusNotFound,
				dtos.ResponseDTO{
					Status:  http.StatusNotFound,
					Message: "Album not found",
					Data:    map[string]interface{}{"data": "Album not found"},
				})
			return
		}

		c.JSON(http.StatusOK, dtos.ResponseDTO{
			Status:  http.StatusOK,
			Message: "Album deleted successfully",
			Data:    map[string]interface{}{"data": "Album deleted successfully"},
		})

	}
}

func GetAllAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var albums []models.Album

		// Fetch all albums
		results, err := albumCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Error fetching albums",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Iterate over results and append to albums
		defer results.Close(ctx)
		for results.Next(ctx) {
			var album models.Album
			err := results.Decode(&album)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{
					Status:  http.StatusInternalServerError,
					Message: "Error Decoding album",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
			albums = append(albums, album)
		}

		c.JSON(http.StatusOK, dtos.ResponseDTO{
			Status:  http.StatusOK,
			Message: "Albums fetched successfully",
			Data:    map[string]interface{}{"data": albums},
		})

	}
}
