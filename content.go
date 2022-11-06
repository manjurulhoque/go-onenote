package onenote

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (client *Client) UpdatePageContent(data []PageContent, url string) (*http.Response, []byte, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, nil, err
	}
	return client.DoRequest(req)
}
