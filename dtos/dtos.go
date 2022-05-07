package dtos

import "web-services-gin/models"

// Package DTOS contains the dto structs for the API.

type ErrorDTO struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type UploadFileDTO struct {
	Message string         `json:"message"`
	Albums  []models.Album `json:"albums"`
}

/*
type UpdateAlbum struct {
	Title  string  `bson:"title,omitempty"`
	Artist string  `bson:"artist,omitempty"`
	Price  float64 `bson:"price,omitempty"`
}
*/
