package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"tiktok-playwright/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type RunActorHandler struct {
	cache map[string]interface{}
}

func NewRunActorHandler() *RunActorHandler {
	return &RunActorHandler{
		cache: make(map[string]interface{}),
	}
}

func (h *RunActorHandler) RunActor(c *gin.Context) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set API token from environment variable
	API_TOKEN := os.Getenv("YOUR_API_TOKEN")

	// Replace input construction with reading from the request body
	input, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	cacheKey, err := json.Marshal(input)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	// Attempt to retrieve cached response from Redis
	cachedResponse, err := utils.GetCache(string(cacheKey)) // Use a suitable cache key
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	if cachedResponse != "" {
		var response interface{}
		err := json.Unmarshal([]byte(cachedResponse), &response)
		if err != nil {
			utils.InternalServerError(c, err)
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}

	// Run the Actor
	resp, err := http.Post("https://api.apify.com/v2/acts/OtzYfK1ndEGdwWFKQ/run-sync-get-dataset-items?token="+API_TOKEN, "application/json", bytes.NewReader(input))
	if err != nil {
		log.Fatalf("Error running actor: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body) // Read the response body into a byte slice
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	// Store the response in Redis cache
	err = utils.SetCache(string(cacheKey), body, 12*time.Hour) // Set expiration as needed
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	var response interface{}
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
