package onenote

import (
	"io"
	"net/http"
)

func (client *Client) ParseResponse(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}

func (client *Client) PageContent(url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	return client.DoRequest(req)
}
