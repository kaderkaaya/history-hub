package model

type EventPage struct {
	Title   string `json:"title"`
	Extract string `json:"extract,omitempty"`
	Image   string `json:"image,omitempty"`
	URL     string `json:"url,omitempty"`
}

type Event struct {
	Year  int         `json:"year"`
	Text  string      `json:"text"`
	Pages []EventPage `json:"pages"`
}
