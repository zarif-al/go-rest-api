package models

// Define Model Struct

type Album struct {
	ID     string  `json:"id" gorm:"primary_key" validate:"required"`
	Title  string  `json:"title" binding:"required" validate:"required"`
	Artist string  `json:"artist" binding:"required" validate:"required"`
	Price  float64 `json:"price" binding:"required" validate:"required"`
}
