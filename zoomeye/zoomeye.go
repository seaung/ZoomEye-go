package zoomeye

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/google/go-querystring/query"
)

var (
	baseURL     = "https://api.zoomeye.org"
	loginAPI    = "https://api.zoomeye.org/%s/search"
	userInfoAPI = "https://api.zoomeye.org/resources-info"
	historyAPI  = "https://api.zoomeye.org/both/search?history=true&ip=%s"
)

type ZoomEyeClient struct {
	mutex       *sync.Mutex
	accessToken string
	apiKey      string
	username    string
	password    string
	Client      *http.Client
}

func NewZoomEyeClient(client *http.Client, token, username, password string) *ZoomEyeClient {
	if client == nil {
		client = http.DefaultClient
	}

	if username == "" && password == "" {
		username = os.Getenv("ZoomEyUsername")
		password = os.Getenv("zoomeyePassword")
	}

	return &ZoomEyeClient{
		apiKey:   token,
		username: username,
		password: password,
		Client:   client,
		mutex:    &sync.Mutex{},
	}
}

func NewEnvZoomEyeClient(client *http.Client) *ZoomEyeClient {
	return NewZoomEyeClient(client, os.Getenv("ZoomEye_API_TOKEN"), "", "")
}

func (z *ZoomEyeClient) NewRequest(method, path string, params interface{}, payloads io.Reader) (*http.Request, error) {
}

func (z *ZoomEyeClient) newRequest(method string, u *url.URL, params interface{}, payloads io.Reader) (*http.Request, error) {
	qs, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	qs.Add("", z.apiKey)
	u.RawQuery = qs.Encode()

	req, err := http.NewRequest(method, u.String(), payloads)
	if err != nil {
		return nil, err
	}

	if payloads != nil {
		req.Header.Add("", fmt.Sprintf("JWT %s", z.accessToken))
	}

	return req, nil
}

func (z *ZoomEyeClient) Login() string {
	return z.accessToken
}
