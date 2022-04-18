package zoomeye

import (
	"encoding/json"
)

var (
	userInfoPath = "/resources-info"
)

type ResourcesInfo struct {
	Plan      string    `json:"plan"`
	Resources Resources `json:"resources"`
	UserInfo  UserInfo  `json:"user_info"`
	QuotaInfo QuotaInfo `json:"quota_info"`
}

type Resources struct {
	Search   int    `json:"search"`
	Stats    string `json:"stats"`
	Interval string `json:"interval"`
}

type UserInfo struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	ExpiredAt string `json:"expired_at"`
}

type QuotaInfo struct {
	RemainFreeQuota  string `json:"remain_free_quota"`
	RemainPayQuota   string `json:"remain_pay_quota"`
	RemainTotalQuota string `json:"remain_total_quota"`
}

func (z *ZoomEyeClient) GetResourcesInfo() *ResourcesInfo {
	var resourceInfos ResourcesInfo

	content, err := z.NewRequest("GET", userInfoPath, nil, nil)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(content, resourceInfos)
	if err != nil {
		return nil
	}

	return &resourceInfos
}
