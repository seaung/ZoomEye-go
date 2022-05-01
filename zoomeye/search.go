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
	Matches []map[string]interface{} `json:"matches"`
	Facets  string                   `json:"facets"`
	Total   int                      `json:"total"`
}

// ty 为0则为host search; ty 为1则为web search;
func (z *ZoomEyeClient) Search(query, facets string, page, ty int) (*Matchers, error) {
	var matches Matchers

	var path string

	if ty == 0 {
		path = fmt.Sprintf(hostSearchPath, query, page, facets)
	} else {
		path = fmt.Sprintf(webSearchPath, query, page, facets)
	}

	resp, err := z.NewRequest("GET", path, nil, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &matches)

	if err != nil {
		return nil, err
	}

	return &matches, nil
}
