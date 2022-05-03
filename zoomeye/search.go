package zoomeye

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Matchers struct {
	Code    int                      `json:"code"`
	Matches []map[string]interface{} `json:"matches"`
	Facets  interface{}              `json:"facets"`
	Total   int                      `json:"total"`
}

func (z *ZoomEyeClient) Search(query, facets string, page int, ty string) (*Matchers, error) {
	client := &http.Client{}
	var matches Matchers
	var path string

	if ty == "web" {
		path = baseURL + fmt.Sprintf(searchAPI, ty, query, page, facets)
	} else {
		path = baseURL + fmt.Sprintf(searchAPI, ty, query, page, facets)
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("API-KEY", z.apiKey)
	req.Header.Set("Authorization", "JWT "+z.accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, err
	}
	return &matches, nil
}
