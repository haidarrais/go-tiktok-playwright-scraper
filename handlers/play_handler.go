package handlers

import (
	"net/http"
	"tiktok-playwright/services"
	"tiktok-playwright/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type PlayHandler struct {
	videoService *services.VideoService
	cache        map[string]string // Cache to store video URLs
}

func NewPlayHandler(videoService *services.VideoService) *PlayHandler {
	return &PlayHandler{
		videoService: videoService,
		cache:        make(map[string]string), // Initialize the cache
	}
}

func (h *PlayHandler) PlayVideo(c *gin.Context) {
	videoURL := c.Query("url")
	if videoURL == "" {
		utils.Error(c, "Video URL is required", http.StatusBadRequest)
		return
	}

	// Attempt to retrieve the cached URL from Redis
	cachedURL, err := utils.GetCache(videoURL)
	if err != nil {
		utils.Error(c, "Error retrieving from cache", http.StatusInternalServerError)
		return
	}

	if cachedURL != "" {
		// If URL is in cache, redirect to cached URL
		c.Redirect(http.StatusFound, cachedURL)
		return
	}

	// If not in cache, proxy the video URL directly to the client
	c.Redirect(http.StatusFound, videoURL)

	// Store the video URL in Redis cache
	err = utils.SetCache(videoURL, videoURL, 24*time.Hour) // Set expiration as needed
	if err != nil {
		utils.Error(c, "Error storing in cache", http.StatusInternalServerError)
	}
}
