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

	"github.com/gin-gonic/gin"
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

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parser := csv.NewReader(file)

	// Skipping the first line
	if _, err := parser.Read(); err != nil {
		panic(err)
	}

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records <- record

	}

}

func uploadAlbums(records chan []string, response chan dtos.UploadFileDTO) {

	defer close(response)
	var uploadFileResponse dtos.UploadFileDTO

	for record := range records {

		album := models.Album{}

		id, err := gonanoid.New()
		if err != nil {
			uploadFileResponse.Albums = append(uploadFileResponse.Albums, album)
			log.Fatal(err)
			break
		}

		album.ID = id
		album.Title = record[0]
		album.Artist = record[1]

		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			uploadFileResponse.Albums = append(uploadFileResponse.Albums, album)
			log.Fatal(err)
			break
		}

		album.Price = price

		response := InsertAlbum(album)

		if response.Error {
			uploadFileResponse.Albums = append(uploadFileResponse.Albums, album)
		}
	}

	if len(uploadFileResponse.Albums) > 0 {
		uploadFileResponse.Message = "Some albums were not uploaded"
	} else {
		uploadFileResponse.Message = "All albums were uploaded"
	}

	response <- uploadFileResponse
}
