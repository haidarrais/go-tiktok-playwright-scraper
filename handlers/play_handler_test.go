package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tiktok-playwright/services"
	"tiktok-playwright/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPlayVideo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	videoService := &services.VideoService{}
	playHandler := NewPlayHandler(videoService)

	router.GET("/play", playHandler.PlayVideo)

	t.Run("Success - Cached URL", func(t *testing.T) {
		utils.SetCache("test_url", "cached_url", 0) // Mock cache
		req, _ := http.NewRequest(http.MethodGet, "/play?url=test_url", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, "cached_url", w.Header().Get("Location"))
	})

	t.Run("Error - Missing URL", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/play", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
