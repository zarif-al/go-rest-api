package controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

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

	go reader(records, saveLocation)

	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
