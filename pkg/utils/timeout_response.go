package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TimeoutResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{
		"code":    http.StatusRequestTimeout,
		"message": "Request timeout",
	})
}
