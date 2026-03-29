package http

import (
	"github.com/gin-gonic/gin"

	Handlers "history-hub/internal/http/handlers"
)

func GinRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) //r.Use olusturunca middle olsuturduk.
	r.Use(gin.Recovery())

	r.GET("/health", Handlers.HealthHandler)
	return r
}
