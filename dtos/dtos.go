package dtos

// Package DTOS contains the dto structs for the API.

type ErrorDTO struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type UploadFileDTO struct {
	Message     string   `json:"message"`
	LineNumbers []string `json:"lineNumbers"`
}
