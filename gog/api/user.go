package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GameResponse struct {
    Products     []int64         `json:"owned"` 
}

func GetUsersGames(bearer string) (*GameResponse, error) {
    request, err := buildRequest("/user/data/games", bearer)
    if err != nil {
        return nil, err
    }

    return sendRequest[GameResponse](request)
}

type GameDetailsResponse struct {
    Title               string          `json:"title"`
    BackgroundImage     string          `json:"backgroundImage"`
    CdKey               string          `json:"cdKey"`
    TextInformation     string          `json:"textInformation"`
    Downloads           [][]interface{} `json:"downloads"`
}

func GetGameDetails(bearer string, gameId int64) (*GameDetailsResponse, error) {
    request, err := buildRequest(fmt.Sprintf("/account/gameDetails/%d.json", gameId), bearer)
    if err != nil {
        return nil, err
    }

    return sendRequest[GameDetailsResponse](request)
}

func buildRequest(path string, bearer string) (*http.Request, error) {
    uri := url.URL {
        Scheme:     "https",
        Host:       "embed.gog.com",
        Path:       path,
    }

    request, err := http.NewRequest("GET", uri.String(), nil)
    if err != nil {
        return nil, err
    }
    request.Header.Set("Bearer", bearer)

    return request, nil
}

func sendRequest[T any](request *http.Request) (*T, error) {
    var responseData T
    client := &http.Client{}

    res, err := client.Do(request)
    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    
    err = json.Unmarshal(body, &responseData)
    if err != nil {
        return nil, err
    }

    return &responseData, nil
}
