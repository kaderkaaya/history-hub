package service

import (
	"context"
	models "history-hub/internal/model"
	provider "history-hub/internal/provider/wikimedia"
)

type EventsService struct {
	client *provider.Client
}

func NewEventsService(client *provider.Client) *EventsService {
	return &EventsService{client: client} //parametre olarak client bekliyor.
}

func (eventService *EventsService) GetEvents(ctx context.Context, lang, typ, month, day string, isToday bool) ([]models.Event, bool, error) {
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
	return events, false, nil
}
