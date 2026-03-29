package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now(),
		"message":   "success",
	})
}
