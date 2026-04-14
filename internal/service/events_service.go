package service

import (
	"context"
	"encoding/json"
	"history-hub/internal/cache"
	models "history-hub/internal/model"
	provider "history-hub/internal/provider/wikimedia"
	"time"
)

type EventsService struct {
	client        *provider.Client
	cache         *cache.RedisClient
	todayTTLHours int
	pastTTLHours  int
}

func NewEventsService(client *provider.Client, redisClient *cache.RedisClient, todayTTL, pastTTL int) *EventsService { //parametre olarak client bekliyor.
	return &EventsService{
		client:        client,
		cache:         redisClient,
		todayTTLHours: todayTTL,
		pastTTLHours:  pastTTL,
	}
}

func (eventService *EventsService) GetEvents(ctx context.Context, lang, typ, month, day string, isToday bool) ([]models.Event, bool, error) {
	key := cache.BuildOnThisDayKey(lang, typ, month, day)

	cachedValue, err := eventService.cache.Get(ctx, key)
	if err == nil {
		var cachedEvents []models.Event
		if json.Unmarshal([]byte(cachedValue), &cachedEvents) == nil {
			return cachedEvents, true, nil
		}
	}

	data, err := eventService.client.GetOnThisDay(lang, typ, month, day)
	if err != nil {
		return nil, false, err
	}

	var events []models.Event
	for _, e := range data.Events {
		var title, url string
		if len(e.Pages) > 0 {
			title = e.Pages[0].Title
			url = e.Pages[0].ContentURLs.Desktop.Page
		}
		events = append(events, models.Event{
			Year: e.Year, Text: e.Text, Title: title, URL: url,
		})
	}

	payload, _ := json.Marshal(events)
	ttl := time.Duration(eventService.pastTTLHours) * time.Hour
	if isToday {
		ttl = time.Duration(eventService.todayTTLHours) * time.Hour
	}
	_ = eventService.cache.Set(ctx, key, string(payload), ttl)

	return events, false, nil
}
