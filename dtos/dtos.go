package dtos

type ResponseDTO struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type UpdateAlbum struct {
	Title  string  `bson:"title,omitempty"`
	Artist string  `bson:"artist,omitempty"`
	Price  float64 `bson:"price,omitempty"`
}
