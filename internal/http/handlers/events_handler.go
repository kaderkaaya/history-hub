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

func (eventHandler *EventsHandler) GetEvents(c *gin.Context) {
	month := c.Query("month")
	day := c.Query("day")
	typ := c.DefaultQuery("type", "events")
	lang := c.DefaultQuery("lang", "en")

	if month == "" {
		c.JSON(400, model.ErrorResponse{Code: 1000, Message: "month is required"})
		return
	}
	if day == "" {
		c.JSON(400, model.ErrorResponse{Code: 1001, Message: "day is required"})
		return
	}

	month = utils.NormalizeMonthDay(month)
	day = utils.NormalizeMonthDay(day)

	if !utils.IsValidType(typ) {
		c.JSON(400, model.ErrorResponse{Code: 1002, Message: "invalid type"})
		return
	}
	if err := utils.ValidateMonthDay(month, day); err != nil {
		c.JSON(400, model.ErrorResponse{Code: 1003, Message: "invalid date"})
		return
	}

	now := time.Now()
	isToday := now.Format("01") == month && now.Format("02") == day
	events, cached, err := eventHandler.service.GetEvents(c.Request.Context(), lang, typ, month, day, isToday)
	if err != nil {
		c.JSON(500, model.ErrorResponse{Code: 1004, Message: "failed to fetch events", Error: err})
		return
	}

	c.JSON(200, model.EventsResponse{
		Date: now.Format("2006") + "-" + month + "-" + day,
		Lang: lang, Type: typ, Cached: cached, Events: events,
	})
}
