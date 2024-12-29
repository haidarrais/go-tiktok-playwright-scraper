package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tiktok-playwright/services"
	"tiktok-playwright/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchVideos(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	videoService := &services.VideoService{}
	searchHandler := NewSearchHandler(videoService)

	router.GET("/search", searchHandler.SearchVideos)

	t.Run("Success - Cached Videos", func(t *testing.T) {
		videos := []services.Video{{Title: "Test Video", URL: "http://test.com", Thumb: "http://thumb.com"}}
		v, _ := json.Marshal(videos)
		utils.SetCache("test_keyword", v, 0) // Mock cache
		req, _ := http.NewRequest(http.MethodGet, "/search?keyword=test_keyword", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
	})

	t.Run("Error - Missing Keyword", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/search", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
