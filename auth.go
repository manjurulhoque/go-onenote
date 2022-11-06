package onenote

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/http"
	"strings"
)

const TokenUrl = "https://login.microsoftonline.com/common/oauth2/v2.0/token"

type RefreshTokenRequest struct {
	ClientId     string `json:"client_id" url:"client_id"`
	Scope        string `json:"scope" url:"scope"`
	RedirectUri  string `json:"redirect_uri" url:"redirect_uri"`
	GrantType    string `json:"grant_type" url:"grant_type"`
	ClientSecret string `json:"client_secret" url:"client_secret"`
	RefreshToken string `json:"refresh_token" url:"refresh_token"`
}

func (refreshTokenRequest *RefreshTokenRequest) toQueryParameters() string {
	value, _ := query.Values(refreshTokenRequest)
	return value.Encode()
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func (client *Client) GetToken() string {
	return ""
}

func (client *Client) GetAccessToken(refreshTokenRequest *RefreshTokenRequest) (*http.Response, *RefreshTokenResponse, error) {
	_, err := json.Marshal(refreshTokenRequest)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", TokenUrl, strings.NewReader(refreshTokenRequest.toQueryParameters()))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, result, err := client.DoRequest(req)
	if err != nil {
		return response, nil, err
	}

	var refreshTokenResponse RefreshTokenResponse
	err = json.Unmarshal(result, &refreshTokenResponse)
	if err != nil {
		return response, &refreshTokenResponse, err
	}

	return response, &refreshTokenResponse, nil
}
