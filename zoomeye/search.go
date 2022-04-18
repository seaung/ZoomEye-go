package zoomeye

import (
	"encoding/json"
	"fmt"
)

var (
	hostSearchPath = "/host/search?query=%s&page=%d&facets=%s"
	webSearchPath  = "/web/search?query=%s&page=%d&facets=%s'"
)

type Matchers struct {
	Matches []interface{} `json:"matches"`
	facets  string        `json:"facets"`
	total   string        `json:"total"`
}

func (z *ZoomEyeClient) HostSearch(query, facets string, page int) *Matchers {
	var matches Matchers

	path := fmt.Sprintf(hostSearchPath, query, page, facets)

	req, err := z.NewRequest("GET", path, nil, nil)

	if err != nil {
		return nil
	}

	err = json.Unmarshal(req, matches)

	if err != nil {
		return nil
	}

	return &matches
}

func (z *ZoomEyeClient) WebSearch(query, facets string, page int) {
	path := fmt.Sprintf(webSearchPath, query, page, facets)
	z.NewRequest("GET", path, nil, nil)
}
