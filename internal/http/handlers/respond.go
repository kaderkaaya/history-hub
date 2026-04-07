package handlers

import (
	model "history-hub/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status, code int, message string, err error) {
	c.JSON(status, model.ErrorResponse{
		Code: code, Message: message, Error: err,
	})
}

func RespondOK(c *gin.Context, payload any) {
	c.JSON(http.StatusOK, payload)
}
