package controllers

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"web-services-gin/dtos"
	"web-services-gin/models"

	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func UploadFile(c *gin.Context) {
	// Current working directory
	dir, _ := os.Getwd()

	file, _ := c.FormFile("file")

	// Retrieve file information
	extension := filepath.Ext(file.Filename)

	if extension != ".csv" {
		c.JSON(400, gin.H{
			"message": "File extension not allowed",
		})
		return
	}

	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	id, err := gonanoid.New()
	if err != nil {
		log.Fatal(err)
	}

	newFileName := id + extension

	saveLocation := dir + "\\uploads\\" + newFileName

	fileSaveErr := c.SaveUploadedFile(file, saveLocation)
	if fileSaveErr != nil {
		log.Fatal(err)
		c.JSON(400, gin.H{"message": "File upload failed"})
	}

	records := make(chan []string)
	response := make(chan dtos.UploadFileDTO)

	go reader(records, saveLocation)

	go uploadAlbums(records, response)

	receiveResponse := <-response
	c.JSON(200, receiveResponse)
}

func reader(records chan []string, saveLocation string) {
	defer close(records)

	file, err := os.Open(saveLocation)
	defer os.Remove(saveLocation)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parser := csv.NewReader(file)

	// Skipping the first line
	if _, err := parser.Read(); err != nil {
		panic(err)
	}

	var lineNumber = 0

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records <- append(record, strconv.Itoa(lineNumber))
		lineNumber++
	}

}

func uploadAlbums(records chan []string, response chan dtos.UploadFileDTO) {

	defer close(response)
	var uploadFileResponse dtos.UploadFileDTO
	var albums []models.Album
	var validate = validator.New()

	var lineNumbers []string

	for record := range records {

		id, err := gonanoid.New()
		if err != nil {
			log.Println(color.Ize(color.Red, "Error : "+err.Error()))
		}

		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Println(color.Ize(color.Red, "Error : "+err.Error()))
		}

		album := models.Album{
			ID:     id,
			Title:  record[0],
			Artist: record[1],
			Price:  price,
		}

		// use validator library to validate required fields
		if validationErr := validate.Struct(&album); validationErr != nil {
			uploadFileResponse.LineNumbers = append(uploadFileResponse.LineNumbers, record[3])
			log.Println(color.Ize(color.Red, "Error : "+validationErr.Error()))
		} else {
			albums = append(albums, album)
			lineNumbers = append(lineNumbers, record[3])
		}

	}

	// TODO: If you want to upload when the albums array reaches a certain length then you have to account for when it can't
	// reach that length.

	result := InsertAlbums(albums)

	if result.Error {
		uploadFileResponse.LineNumbers = append(uploadFileResponse.LineNumbers, lineNumbers...)
		log.Println(color.Ize(color.Red, result.Message))
	}

	if len(uploadFileResponse.LineNumbers) > 0 {
		uploadFileResponse.Message = "Some albums were not uploaded"
	} else {
		uploadFileResponse.Message = "All albums were uploaded"
	}

	response <- uploadFileResponse
}
