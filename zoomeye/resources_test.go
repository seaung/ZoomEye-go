package zoomeye

import (
	"net/url"
	"testing"
)

func TestGetResource(t *testing.T) {
	z := NewEnvZoomEyeClient()

	result, err := z.GetResourcesInfo()
	if err != nil {
		t.Logf("Can't Get resource ifnormation from GetResourcesInfo : %v\n", err)
	}

	t.Logf("resource info : %v\n", result)
}

func TestGetResourceWithPath(t *testing.T) {
	z := NewEnvZoomEyeClient()

	t.Logf("current url : %v\n", baseURL+userInfoPath)

	u, err := url.Parse(baseURL + userInfoPath)
	if err != nil {
		t.Logf("Can't parse taget url in TestGetResourceWithPath : %v\n", err)
	}

	result, err := z.newRequest("GET", u, nil, nil)
	if err != nil {
		t.Logf("Can't get resource info content : %v\n", err)
	}

	t.Logf("GetResourcesInfo content : %v\n", string(result))
}
