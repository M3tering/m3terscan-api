package main

import (
	"log"
	"m3terscan-api/internal/api"
	"m3terscan-api/internal/blockchain"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	client, err := blockchain.GetClient()
	if err != nil {
		log.Fatal("Failed to initialize client:", err)
	}
	defer client.Close()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://ap-dashboard-kappa.vercel.app", "https://m3terscan-rr.vercel.app", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to m3terscan api 😎",
		})
	})
	router.GET("/m3ter/:id/daily", api.GetDaily)
	router.GET("/m3ter/:id/weekly", api.GetWeekly)
	router.GET("/m3ter/:id/monthly", func(ctx *gin.Context) {
		api.GetMonthly(ctx, client)
	})
	router.GET("/proposal", func(ctx *gin.Context) {
		api.GetCommitState(ctx, client)
	})
	router.GET("m3ter/:id/activities", api.GetActivities)

	router.Run() // listens on 0.0.0.0:8080 by default
}
