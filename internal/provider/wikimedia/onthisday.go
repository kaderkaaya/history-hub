package wikimedia

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OnThisDayResponse struct { //wikimedianın dondugu response.
	Events []struct {
		Year  int    `json:"year"`
		Text  string `json:"text"`
		Pages []struct {
			Title       string `json:"title"`
			ContentURLs struct {
				Desktop struct {
					Page string `json:"page"`
				} `json:"desktop"`
			} `json:"content_urls"`
		} `json:"pages"`
	} `json:"events"`
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
