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

type Data struct {
	Component   []string `json:"component"`
	DB          []string `json:"db"`
	Description []string `json:"description"`
	Domains     []string `json:"domains"`
	Framework   []string `json:"framework"`
	Geoinfo     Geoinfo  `json:"geoinfo"`
	Headers     string   `json:"headers"`
	IP          []string `json:"ip"`
	Keyword     string   `json:"keyword"`
	Language    []string `json:"language"`
	Server      []Server `json:"server"`
	Site        string   `json:"site"`
	System      []System `json:"system"`
	Timestamp   string   `json:"timestamp"`
	Title       string   `json:"title"`
	Waf         []string `json:"waf"`
	Webapp      []string `json:"webapp"`
}

type Geoinfo struct {
	PoweredBy      string       `json:"poweredBy"`
	Asn            interface{}  `json:"asn"`
	Aso            interface{}  `json:"aso"`
	BaseStation    string       `json:"base_station"`
	City           City         `json:"city"`
	Continent      Continent    `json:"continent"`
	Country        Country      `json:"country"`
	Idc            string       `json:"idc"`
	Isp            string       `json:"isp"`
	Location       interface{}  `json:"location"`
	Organization   string       `json:"organization"`
	OrganizationCN string       `json:"organization_CN"`
	Subdivisions   Subdivisions `json:"subdivisions"`
}

type Server struct {
	Chinese string `json:"chinese"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type System struct {
	Chinese string      `json:"chinese"`
	Distrib interface{} `json:"distrib"`
	Name    string      `json:"name"`
	Release interface{} `json:"release"`
	Version interface{} `json:"version"`
}

type City struct {
	GeonameId string `json:"geoname_id"`
	Names     Names  `json:"names"`
}

type Names struct {
	En   string `json:"en"`
	ZhCN string `json:"zh-CN"`
}

type Continent struct {
	Code      int         `json:"code"`
	GeonameId interface{} `json:"geoname_id"`
	Names     Names       `json:"names"`
}

type Country struct {
	Code      interface{} `json:"code"`
	GeonameId interface{} `json:"geoname_id"`
	Names     Names       `json:"name"`
}

type Subdivisions struct {
	Code      interface{} `json:"code"`
	GeonameId interface{} `json:"geoname_id"`
	Names     Names       `json:"names"`
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
