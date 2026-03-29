package handlers

import (
	"github.com/gin-gonic/gin"

	service "history-hub/internal/service"
)

func GetEventsHandler(ctx *gin.Context) {
	events, err := service.GetEvents()

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, events)

}
