package http

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	handlers "history-hub/internal/http/handlers"
	utils "history-hub/pkg/utils"
)

func HistoryHubRouter(eventsHandler *handlers.EventsHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) //r.Use olusturunca middle olsuturduk.
	r.Use(gin.Recovery())
	r.Use(timeout.New(
		timeout.WithTimeout(10*time.Second), //burda eğer bir endpointin süresi verdiğimiz süreden fazla sürerse timeout hatasi verir.
		timeout.WithResponse(utils.TimeoutResponse),
	))

	r.GET("/health", handlers.HealthHandler)
	events := r.Group("/events")
	events.GET("/get-events", eventsHandler.GetEvents)
	return r
}
