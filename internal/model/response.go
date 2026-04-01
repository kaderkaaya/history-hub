package model

type EventsResponse struct {
	Date   string  `json:"date"`
	Lang   string  `json:"lang"`
	Type   string  `json:"type"`
	Cached bool    `json:"cached"`
	Events []Event `json:"events"`
}
