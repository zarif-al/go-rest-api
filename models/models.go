package models

// Define Model Struct

type Album struct {
	ID     string  `json:"id" gorm:"primary_key"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}
