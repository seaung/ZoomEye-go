package zoomeye

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	domainPath = "/domain/search?q=%s&type=%d&page=%d"
)

type DomainInfo struct {
	Status int           `json:"status"`
	List   []interface{} `json:"list"`
	Total  int           `json:"total"`
	Msg    string        `json:"msg"`
	Type   int           `json:"type"`
}

type List struct {
	Name      string   `json:"name"`
	Timestamp string   `json:"timestamp"`
	Ip        []string `json:"ip"`
}

func (z *ZoomEyeClient) DomainSearch(query string, types, page int) (*DomainInfo, error) {
	var domainInfo DomainInfo
	client := &http.Client{}

	path := baseURL + fmt.Sprintf(domainPath, query, types, page)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("API-KEY", z.apiKey)
	req.Header.Set("Authorization", "JWT "+z.accessToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&domainInfo); err != nil {
		return nil, err
	}
	return &domainInfo, nil
}
