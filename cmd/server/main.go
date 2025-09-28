package main

import (
	"m3terscan-api/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/m3ter/:id/daily", api.GetDaily)
	router.GET("/m3ter/:id/weekly", api.GetWeekly)

	router.Run() // listens on 0.0.0.0:8080 by default
}
