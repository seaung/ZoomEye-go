package zoomeye

import (
	"encoding/json"
	"net/http"
)

var (
	userInfoPath = "/resources-info"
)

/*
{
	"code": 60000,
	"plan": "developer",
	"resources": {"search": 80, "stats": 100, "interval": "month"},
	"user_info": {"name": "", "role": "developer", "expired_at": ""},
	"quota_info": {"remain_free_quota": 0, "remain_pay_quota": 0, "remain_total_quota": 9}
}
*/

type ResourcesInfo struct {
	Code      int       `json:"code"`
	Plan      string    `json:"plan"`
	Resources Resources `json:"resources"`
	UserInfo  UserInfo  `json:"user_info"`
	QuotaInfo QuotaInfo `json:"quota_info"`
}

type Resources struct {
	Search   int    `json:"search"`
	Stats    int    `json:"stats"`
	Interval string `json:"interval"`
}

type UserInfo struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	ExpiredAt string `json:"expired_at"`
}

type QuotaInfo struct {
	RemainFreeQuota  int `json:"remain_free_quota"`
	RemainPayQuota   int `json:"remain_pay_quota"`
	RemainTotalQuota int `json:"remain_total_quota"`
}

func (z *ZoomEyeClient) GetResourcesInfo() (*ResourcesInfo, error) {
	client := &http.Client{}
	path := baseURL + userInfoPath
	var resourceInfos ResourcesInfo

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

	if err = json.NewDecoder(resp.Body).Decode(&resourceInfos); err != nil {
		return nil, err
	}

	return &resourceInfos, nil
}
