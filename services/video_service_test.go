package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchVideos(t *testing.T) {
	videoService := NewVideoService()

	t.Run("Fetch Videos - Valid Keyword", func(t *testing.T) {
		videos, err := videoService.FetchVideos("test")
		assert.NoError(t, err)
		assert.NotEmpty(t, videos)
	})

	t.Run("Fetch Videos - Invalid Keyword", func(t *testing.T) {
		// Assuming the service handles invalid keywords gracefully
		videos, err := videoService.FetchVideos("invalid_keyword")
		assert.NoError(t, err)
		assert.Empty(t, videos)
	})
}
