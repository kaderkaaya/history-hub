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
	
	// Determine which event list to use based on what is populated
	var sourceEvents []provider.WikimediaEvent
	if len(data.Events) > 0 {
		sourceEvents = data.Events
	} else if len(data.Deaths) > 0 {
		sourceEvents = data.Deaths
	} else if len(data.Holidays) > 0 {
		sourceEvents = data.Holidays
	} else if len(data.Selected) > 0 {
		sourceEvents = data.Selected
	} else if len(data.Births) > 0 {
		sourceEvents = data.Births
	}

	for _, e := range sourceEvents {
		pages := make([]models.EventPage, 0, len(e.Pages))
		for _, p := range e.Pages {
			image := p.OriginalImage.Source
			if image == "" {
				image = p.Thumbnail.Source
			}
			title := p.Titles.Normalized
			if title == "" {
				title = p.Title
			}
			pages = append(pages, models.EventPage{
				Title:   title,
				Extract: p.Extract,
				Image:   image,
				URL:     p.ContentURLs.Desktop.Page,
			})
		}
		events = append(events, models.Event{
			Year:  e.Year,
			Text:  e.Text,
			Pages: pages,
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
