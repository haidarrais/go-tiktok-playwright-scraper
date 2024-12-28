// main.go
package main

import (
	"log"
	"tiktok-playwright/handlers"
	"tiktok-playwright/services"
	"tiktok-playwright/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis
	utils.InitializeRedis("localhost:6379", "", 0)

	r := gin.Default()

	// Initialize services
	videoService := services.NewVideoService()

	searchHandler := handlers.NewSearchHandler(videoService)
	playHandler := handlers.NewPlayHandler(videoService)

	// Define routes
	r.GET("/search", searchHandler.SearchVideos)
	r.GET("/play", playHandler.PlayVideo)

	port := ":3030"
	log.Printf("Server running on %s", port)
	r.Run(port)
}
