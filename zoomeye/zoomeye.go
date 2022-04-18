package zoomeye

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

var (
	baseURL    = "https://api.zoomeye.org"
	historyAPI = "/both/search?history=true&ip=%s"
)

type ZoomEyeClient struct {
	apiKey string
}

func NewZoomEyeClient(apikey string) *ZoomEyeClient {
	return &ZoomEyeClient{
		apiKey: apikey,
	}
}

func NewEnvZoomEyeClient() *ZoomEyeClient {
	return &ZoomEyeClient{
		apiKey: os.Getenv("ZOOMEYE_API_KEY"),
	}
}

func (z *ZoomEyeClient) newRequest(method string, u *url.URL, params interface{}, payloads io.Reader) ([]byte, error) {
	client := &http.Client{}
	qs, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	u.RawQuery = qs.Encode()

	req, err := http.NewRequest(method, u.String(), payloads)

	if err != nil {
		return nil, err
	}

	req.Header.Set("API-KEY", z.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp != nil && resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (z *ZoomEyeClient) NewRequest(method string, path string, params interface{}, payloads io.Reader) ([]byte, error) {
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, err
	}

	return z.newRequest(method, u, params, payloads)
}

func (z *ZoomEyeClient) GetHistory(ipaddr string) ([]byte, error) {
	path := fmt.Sprintf(historyAPI, ipaddr)

	req, err := z.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
