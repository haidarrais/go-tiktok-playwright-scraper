package handlers

import (
	"encoding/json"
	"net/http"
	"tiktok-playwright/services"
	"tiktok-playwright/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	videoService *services.VideoService
	cache        map[string]interface{}
}

func NewSearchHandler(videoService *services.VideoService) *SearchHandler {
	return &SearchHandler{
		videoService: videoService,
		cache:        make(map[string]interface{}),
	}
}

func (h *SearchHandler) SearchVideos(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		utils.Error(c, "Keyword is required", http.StatusBadRequest)
		return
	}

	// Attempt to retrieve cached videos from Redis
	cachedVideos, err := utils.GetCache(keyword)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}
	if cachedVideos != "" {
		var cVideos []services.Video
		err := json.Unmarshal([]byte(cachedVideos), &cVideos)
		if err != nil {
			panic(err)
		}
		utils.Success(c, cVideos)
		return
	}

	videos, err := h.videoService.FetchVideos(keyword)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	v, err := json.Marshal(&videos)
	if err != nil {
		panic(err)
	}

	// Store the fetched videos in Redis cache
	err = utils.SetCache(keyword, v, 12*time.Hour) // Set expiration as needed
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, videos)
}
