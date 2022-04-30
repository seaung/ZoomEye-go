package zoomeye

import (
	"net/url"
	"os"
	"testing"
)

func TestNewZoomEyeClient(t *testing.T) {
	z := NewZoomEyeClient(os.Getenv("ZOOMEYE_API_KEY"))
	t.Logf("New zoomeye client : %v\n", z.apiKey)
}

func TestNewEnvZoomEyeClient(t *testing.T) {
	z := NewEnvZoomEyeClient()
	t.Logf("New env zoomeye client : %v\n", z.apiKey)
}

func TestRequest(t *testing.T) {
	z := NewEnvZoomEyeClient()

	u, err := url.Parse("https://www.baidu.com/")
	if err != nil {
		t.Logf("Can't parse url : %v\n", err)
	}

	content, err := z.newRequest("GET", u, nil, nil)
	if err != nil {
		t.Logf("Can't get response content : %v\n", err)
	}

	t.Logf("response content : %v\n", content)
}

func TestNewRequest(t *testing.T) {
	z := NewEnvZoomEyeClient()

	content, err := z.NewRequest("GET", "", nil, nil)
	if err != nil {
		t.Logf("Get response from NewRequest method : %v\n", err)
	}
	t.Logf("Get response from NewRequest method : %v\n", content)
}
