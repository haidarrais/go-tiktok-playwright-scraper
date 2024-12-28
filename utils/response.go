package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success sends a JSON response with a success message and data
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// Error sends a JSON response with an error message
func Error(c *gin.Context, message string, code int) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "error",
		"message": "Resource not found",
	})
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": "Internal server error",
		"error":   err.Error(),
	})
}
