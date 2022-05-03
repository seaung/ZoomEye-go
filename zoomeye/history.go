package zoomeye

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	historyPath = "/both/search?history=true&ip=%s"
)

type HistoryRecord struct {
	Count int           `json:"count"`
	Data  []interface{} `json:"data"`
}

func (z *ZoomEyeClient) BothHistoty(ipaddr string) (hr *HistoryRecord, err error) {
	client := &http.Client{}
	path := baseURL + fmt.Sprintf(historyPath, ipaddr)

	var records HistoryRecord

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

	if err = json.NewDecoder(resp.Body).Decode(&records); err != nil {
		return nil, err
	}

	return &records, nil
}
