package model

type Event struct {
	Year  int    `json:"year"`
	Text  string `json:"text"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
