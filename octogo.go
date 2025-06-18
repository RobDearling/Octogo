package octogo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseUrl = "https://api.octopus.energy/"
	userAgent      = "octogo/" + libraryVersion
)

type Client struct {
	HttpClient *http.Client
	BaseUrl    *url.URL
	Auth       string
	UserAgent  string

	Meter MeterService
}

type Response struct {
	*http.Response
}

func NewClient(apiKey string) *Client {
	httpClient := &http.Client{}
	baseUrl, _ := url.Parse(defaultBaseUrl)

	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))

	c := &Client{HttpClient: httpClient, BaseUrl: baseUrl, Auth: auth, UserAgent: userAgent}

	c.Meter = NewMeterService(c)

	return c
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Basic "+c.Auth)
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &Response{Response: resp}

	if resp.StatusCode != http.StatusOK {
		return response, nil
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return response, err
}
