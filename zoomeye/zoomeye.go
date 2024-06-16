package zoomeye

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	baseURL       = "https://api.zoomeye.org"
	historyAPI    = "/both/search?history=true&ip=%s"
	searchHostAPI = "/host/search/?query=%s&page=%d&facets=%s"
	searchWebAPI  = "/web/search/?query=%s&page=%d&facets=%s"
	loginPath     = "/user/login"
)

type ZoomEyeClient struct {
	username    string
	password    string
	accessToken string
	apiKey      string
}

type Token struct {
	AccessToken string `json:"access_token"`
}

func NewZoomEyeWithAuth(accessToken string) *ZoomEyeClient {
    return &ZoomEyeClient{
        accessToken: accessToken,
    }
}

func NewZoomEyeClient(username, password, accessToken, apikey string) *ZoomEyeClient {
	return &ZoomEyeClient{
		apiKey:      apikey,
		username:    username,
		password:    password,
		accessToken: accessToken,
	}
}

func NewEnvZoomEyeClient() *ZoomEyeClient {
	return &ZoomEyeClient{
		apiKey:   os.Getenv("ZOOMEYE_API_KEY"),
		username: os.Getenv("ZOOMYEY_USERNAME"),
		password: os.Getenv("ZOOMYEY_PASSWORD"),
	}
}

func (z *ZoomEyeClient) Login() (string, error) {
	client := &http.Client{}
	path := baseURL + loginPath
	var token Token

	params := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, z.username, z.password)

	req, err := http.NewRequest("POST", path, strings.NewReader(params))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	z.accessToken = token.AccessToken

	return token.AccessToken, nil
}
