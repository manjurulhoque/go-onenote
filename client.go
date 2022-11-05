package onenote

import (
	"fmt"
	"io"
	"net/http"
)

const DefaultRestUrl = "https://graph.microsoft.com/v1.0"

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

func (client *Client) newRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", DefaultRestUrl, path), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (client *Client) doRequest(req *http.Request) (*http.Response, []byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.ApiKey))
	response, err := client.HttpClient.Do(req)
	if err != nil {
		return response, nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return response, nil, err
	}
	return response, body, err
}
