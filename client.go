package onenote

import (
	"fmt"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	Host       string
	Base       string
}

func NewClient(apiKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
	}
}

func (client *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.ApiKey))
	res, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, err
}
