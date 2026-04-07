package handlers

import (
	"time"

	"github.com/gin-gonic/gin"

	model "history-hub/internal/model"
	service "history-hub/internal/service"
	utils "history-hub/pkg/utils"
)

type EventsHandler struct {
	service *service.EventsService
}

func NewEventsHandler(service *service.EventsService) *EventsHandler {
	return &EventsHandler{service: service}
}

func (eventHandler *EventsHandler) GetTodayEvents(c *gin.Context) {
	req := model.GetTodayRequest{Typ: "events", Lang: "en"}

	if err := c.ShouldBindQuery(&req); err != nil {
		RespondError(c, 400, 1001, "Validation Error", err)
		return
	}
	now := time.Now()
	isToday := true
	month := now.Month()
	day := now.Day()
	events, cached, err := eventHandler.service.GetEvents(c.Request.Context(), req.Lang, req.Typ, utils.NormalizeMonthDayInt(int(month)), utils.NormalizeMonthDayInt(day), isToday)
	if err != nil {
		RespondError(c, 500, 1004, "failed to fetch events", err)
		return
	}

	c.JSON(200, model.EventsResponse{
		Date: now.Format("2006-01-02"),
		Lang: req.Lang, Type: req.Typ, Cached: cached, Events: events,
	})
}

func (eventHandler *EventsHandler) GetEvents(c *gin.Context) {
	req := model.GetEventsRequest{Typ: "events", Lang: "en"}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, model.ErrorResponse{Code: 1001, Message: "Validation Error", Error: err})
		return
	}
	now := time.Now()
	isToday := int(now.Month()) == req.Month && int(now.Day()) == req.Day
	events, cached, err := eventHandler.service.GetEvents(c.Request.Context(), req.Lang, req.Typ, utils.NormalizeMonthDayInt(req.Month), utils.NormalizeMonthDayInt(req.Day), isToday)
	if err != nil {
		c.JSON(500, model.ErrorResponse{Code: 1004, Message: "failed to fetch events", Error: err})
		return
	}

	c.JSON(200, model.EventsResponse{
		Date: now.Format("2006-01-02"),
		Lang: req.Lang, Type: req.Typ, Cached: cached, Events: events,
	})
}
