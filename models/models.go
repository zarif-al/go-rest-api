package models

// Define Model Struct

type Album struct {
	ID     string  `json:"id" gorm:"primary_key"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
