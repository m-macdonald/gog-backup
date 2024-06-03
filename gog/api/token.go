package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TokenResponse struct {
    ExpiresIn       int     `json:"expires_in"`
    Scope           string  `json:"scope"`
    TokenType       string  `json:"token_type"`
    AccessToken     string  `json:"access_token"`
    UserId          string  `json:"user_id"`
    RefreshToken    string  `json:"refresh_token"`
    SessionId       string  `json:"session_id"`
}

func httpClient() *http.Client {
    return &http.Client {
        
    }
}

func GetToken(grantType string, code string, redirectUri string, refreshToken string) (TokenResponse, error) {
    var tokenResponse TokenResponse
    var uri = url.URL {
        Scheme:     "https",
        Host:       "auth.gog.com",
        Path:       "token",
    }

    queryParams := url.Values {}
    queryParams.Add("client_id", CLIENT_ID)
    queryParams.Add("client_secret", CLIENT_SECRET)
    queryParams.Add("grant_type", grantType)
    queryParams.Add("code", code)
    queryParams.Add("redirect_uri", redirectUri)
    queryParams.Add("refresh_token", refreshToken)

    uri.RawQuery = queryParams.Encode()

    res, err := httpClient().Get(uri.String())
    if err != nil {
        fmt.Println(err)
        return tokenResponse, err
    }

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return tokenResponse, err
    }

    json.Unmarshal(body, &tokenResponse)

    println(string(body[:]))

    return tokenResponse, nil
}
