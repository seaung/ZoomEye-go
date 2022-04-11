package zoomeye

import (
	"net/http"
	"os"
)

var (
	baseURL = "https://api.zoomeye.org"
)

type ZoomEyeClient struct {
	accessToken string
	Client      *http.Client
}

func NewZoomEyeClient(client *http.Client, token string) *ZoomEyeClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &ZoomEyeClient{
		accessToken: token,
		Client:      client,
	}
}

func NewEnvZoomEyeClient(client *http.Client) *ZoomEyeClient {
	return NewZoomEyeClient(client, os.Getenv("ZoomEye_API_TOKEN"))
}

func (z *ZoomEyeClient) NewRequest(method string) {}
