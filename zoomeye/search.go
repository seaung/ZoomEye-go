package zoomeye

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Hosts struct {
	Code      int       `json:"code"`
	Total     int       `json:"total"`
	Available int       `json:"available"`
	Matches   []Matches `json:"matches"`
	//Facets    interface{}               `json:"facets"`
}

type Matches struct {
	Jarm      string           `json:"jarm"`
	Ico       Ico              `json:"ico"`
	Txtfile   Txtfile          `json:"txtfile"`
	IP        string           `json:"ip"`
	PortInfo  PortInfo         `json:"portinfo"`
	Timestamp string           `json:"timestamp"`
	Geoinfo   Geoinfo          `json:"geoinfo"`
	Protocol  Protocol         `json:"protocol"`
	Honeypot  int              `json:"honeypot"`
	Whois     map[string]Whois `json:"whois"`
}

type Ico struct {
	Mmh3 string `json:"mmh3"`
	Md5  string `json:"md5"`
}

type Txtfile struct {
	Robotsmd5   string `json:"robotsmd5"`
	Securitymd5 string `json:"securitymd5"`
}

/*
"version": "",
"device": "",
"extrainfo": "",
"rdns": "",
"app": "nginx",
"banner"
*/
type PortInfo struct {
	Hostname  string   `json:"hostname"`
	Os        string   `json:"os"`
	Port      int      `json:"port"`
	Service   string   `json:"service"`
	Title     []string `json:"title"`
	Version   string   `json:"version"`
	Device    string   `json:"device"`
	Extrainfo string   `json:"extrainfo"`
	Rdns      string   `json:"rdns"`
	App       string   `json:"app"`
	Banner    string   `json:"banner"`
}

type Geoinfo struct {
	Continent      Continent    `json:"continent"`
	Country        Country      `json:"country"`
	BaseStation    string       `json:"base_station"`
	City           City         `json:"city"`
	Isp            string       `json:"isp"`
	Organization   string       `json:"organization"`
	Idc            string       `json:"idc"`
	Location       Location     `json:"location"`
	Aso            interface{}  `json:"aso"`
	Asn            string       `json:"asn"`
	Subdivisions   Subdivisions `json:"subdivisions"`
	PowerdBy       string       `json:"PowerdBy"`
	Scene          Scene        `json:"scene"`
	OrganizationCN string       `json:"organization_CN"`
}

/*
"continent": {
	"code": "AP",
	"names": {
		"en": "Asia",
		"zh-CN": "亚洲"
	},
	"geoname_id": null
},
*/
type Continent struct {
	Code      string      `json:"code"`
	Names     Names       `json:"names"`
	GeonameId interface{} `json:"geoname_id"`
}

/*
"country": {
	"code": "CN",
	"names": {
		"en": "China",
		"zh-CN": "中国"
	},
	"geoname_id": null
},
*/
type Country struct {
	Code      string      `json:"code"`
	Names     Names       `json:"names"`
	GeonameId interface{} `json:"geoname_id"`
}

type Names struct {
	En   string `json:"en"`
	ZhCN string `json:"zh-CN"`
}

/*
"city": {
	"names": {
		"en": "Beijing",
		"zh-CN": "北京"
	},
	"geoname_id": null
},
*/
type City struct {
	Names     Names       `json:"name"`
	GeonameId interface{} `json:"geoname_id"`
}

/*
"subdivisions": {
	"names": {
		"en": "Beijing",
		"zh-CN": "北京"
	},
	"geoname_id": null
},
*/
type Subdivisions struct {
	Names     Names       `json:"names"`
	GeonameId interface{} `json:"geoname_id"`
}

type Location struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}

type Scene struct {
	En string `json:"en"`
	Cn string `json:"cn"`
}

type Protocol struct {
	Application string `json:"application"`
	Probe       string `json:"probe"`
	Transport   string `json:"transport"`
}

type Whois struct {
	AdminC       string       `json:"admin_c"`
	Country      string       `json:"country"`
	Descr        string       `json:"descr"`
	Inetnum      string       `json:"inetnum"`
	IpEnd        string       `json:"ip_end"`
	Irt          []Irt        `json:"irt"`
	LastModified string       `json:"last_modified"`
	MntBy        string       `json:"mnt_by"`
	MntIrt       string       `json:"mnt_irt"`
	MntLower     string       `json:"mnt_lower"`
	MntRoutes    string       `json:"mnt_routes"`
	NetName      string       `json:"net_name"`
	Notify       string       `json:"notify"`
	Organization Organization `json:"organization"`
	Person       []Person     `json:"person"`
	Role         []Role       `json:"role"`
	Source       string       `json:"source"`
	Status       string       `json:"status"`
	TechC        string       `json:"tech_c"`
}

type Irt struct {
	AbusMainbox  string `json:"abus_maxinbox"`
	Address      string `json:"address"`
	AdminC       string `json:"admin_c"`
	Auth         string `json:"auth"`
	Email        string `json:"email"`
	FaxNo        string `json:"fax_no"`
	Irt          string `json:"irt"`
	IrtNfy       string `json:"irt_nfy"`
	LastModified string `json:"last_modified"`
	MntBy        string `json:"mnt_by"`
	Notify       string `json:"notify"`
	Phone        string `json:"phone"`
	Source       string `json:"source"`
	TechC        string `json:"tech_c"`
}

type Organization struct {
	Address      string `json:"address"`
	AdminC       string `json:"admin_c"`
	Country      string `json:"country"`
	Descr        string `json:"descr"`
	Email        string `json:"email"`
	FaxNo        string `json:"fax_no"`
	LastModified string `json:"last_modified"`
	MntBy        string `json:"mnt_by"`
	MntRef       string `json:"mnt_ref"`
	Notify       string `json:"notify"`
	Org          string `json:"org"`
	OrgName      string `json:"org_name"`
	Organization string `json:"organization"`
	Phone        string `json:"phone"`
	Source       string `json:"source"`
	TechC        string `json:"tech_c"`
}

type Person struct {
	AbusMainbox  string `json:"abus_maxinbox"`
	Address      string `json:"address"`
	Country      string `json:"country"`
	Email        string `json:"email"`
	FaxNo        string `json:"fax_no"`
	LastModified string `json:"last_modified"`
	MntBy        string `json:"mnt_by"`
	NicHdl       string `json:"nic_hdl"`
	Notify       string `json:"notify"`
	Person       string `json:"person"`
	Phone        string `json:"phone"`
	Source       string `json:"source"`
}

type Role struct {
	AbusMainbox  string `json:"abus_maxinbox"`
	Address      string `json:"address"`
	AdminC       string `json:"admin_c"`
	Country      string `json:"country"`
	Email        string `json:"email"`
	FaxNo        string `json:"fax_no"`
	LastModified string `json:"last_modified"`
	MntBy        string `json:"mnt_by"`
	NicHdl       string `json:"nic_hdl"`
	Notify       string `json:"notify"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
	Source       string `json:"source"`
	TechC        string `json:"tech_c"`
}

func (z *ZoomEyeClient) WebSearch(query, facets string, page int) (*Hosts, error) {
	var hosts Hosts
	return &hosts, nil
}

func (z *ZoomEyeClient) SeachHosts(url, accessToken string) (*Hosts, error) {
	var matches Hosts
    client := &http.Client{}

    request, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    request.Header.Set("Authorization", "JWT " + accessToken)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, err
	}
	return &matches, nil
}

func (z *ZoomEyeClient) HostSearch(query, facets string, page int) (*Hosts, error) {
	client := &http.Client{}
	var matches Hosts
	var path string

	path = baseURL + fmt.Sprintf(searchHostAPI, query, page, facets)

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

