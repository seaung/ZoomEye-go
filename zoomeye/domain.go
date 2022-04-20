package zoomeye

import (
	"encoding/json"
	"fmt"
)

var (
	domainPath = "/domain/search?q=%s&type=%d&page=%d"
)

type DomainInfo struct {
	Status string `json:"status"`
	List   []List `json:"list"`
	Total  int    `json:"total"`
	Msg    string `json:"msg"`
	Type   int    `json:"type"`
}

type List struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Ip        string `json:"ip"`
}

func (z *ZoomEyeClient) DomainSearch(query string, types, page int) (*DomainInfo, error) {
	var domainInfo DomainInfo

	path := fmt.Sprintf(domainPath, query, types, page)

	content, err := z.NewRequest("GET", path, nil, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, domainInfo)
	if err != nil {
		return nil, err
	}

	return &domainInfo, nil
}
