package controllers

import (
	"net/http"
	"time"
	dtos "web-services-gin/dtos"

	"github.com/gin-gonic/gin"
)

type UrlType struct {
	Url string `json:"url"`
}

func CheckStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url UrlType

		if err := c.BindJSON(&url); err != nil {
			c.JSON(http.StatusBadRequest, dtos.ResponseDTO{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		client := http.Client{
			Timeout: 2 * time.Second,
		}

		resp, err := client.Get(url.Url)

		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Url Status",
				Data:    map[string]interface{}{"status": 408},
			})
			return
		}

		// Return album
		c.JSON(http.StatusOK, dtos.ResponseDTO{
			Status:  http.StatusOK,
			Message: "Url Status",
			Data:    map[string]interface{}{"status": resp.StatusCode},
		})
	}
}
