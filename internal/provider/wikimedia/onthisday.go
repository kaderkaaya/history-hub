package wikimedia

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WikimediaImage struct {
	Source string `json:"source"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type WikimediaTitles struct {
	Canonical  string `json:"canonical"`
	Normalized string `json:"normalized"`
	Display    string `json:"display"`
}

type WikimediaPage struct {
	Title         string          `json:"title"`
	Titles        WikimediaTitles `json:"titles"`
	Extract       string          `json:"extract"`
	Thumbnail     WikimediaImage  `json:"thumbnail"`
	OriginalImage WikimediaImage  `json:"originalimage"`
	ContentURLs   struct {
		Desktop struct {
			Page string `json:"page"`
		} `json:"desktop"`
	} `json:"content_urls"`
}

type WikimediaEvent struct {
	Year  int             `json:"year"`
	Text  string          `json:"text"`
	Pages []WikimediaPage `json:"pages"`
}

type OnThisDayResponse struct { //wikimedianın dondugu response.
	Events   []WikimediaEvent `json:"events,omitempty"`
	Deaths   []WikimediaEvent `json:"deaths,omitempty"`
	Holidays []WikimediaEvent `json:"holidays,omitempty"`
	Selected []WikimediaEvent `json:"selected,omitempty"`
	Births   []WikimediaEvent `json:"births,omitempty"`
}

func (c *Client) GetOnThisDay(lang, typ, month, day string) (*OnThisDayResponse, error) {
	url := fmt.Sprintf("%s/feed/v1/wikipedia/%s/onthisday/%s/%s/%s",
		c.WikimediaBaseURL, lang, typ, month, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data OnThisDayResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
