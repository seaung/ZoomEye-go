package zoomeye

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	baseURL    = "https://api.zoomeye.org"
	historyAPI = "/both/search?history=true&ip=%s"
	searchAPI  = "/%s/search/?query=%s&page=%d&facets=%s"
	loginPath  = "/user/login"
)

type ZoomEyeClient struct {
	username    string
	password    string
	accessToken string
	apiKey      string
}

type Token struct {
	accessToken string `json:"access_token"`
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
		username: os.Getenv("ZOOMEYE_USERNAME"),
		password: os.Getenv("ZOOMEYE_PASSWORD"),
	}
}

func (z *ZoomEyeClient) Login() (string, error) {

	var token Token

	params := fmt.Sprintf(`{{"username": "%s", "password": "%s"}}`, z.username, z.password)

	resp, err := http.Post(baseURL+loginPath, "application/json", strings.NewReader(params))
	if err != nil {
		return "", err
	}

	if resp == nil && resp.StatusCode != http.StatusOK {
		return "", err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "Can't get access token", err
	}
	return token.accessToken, nil
}
