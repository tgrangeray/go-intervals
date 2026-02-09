package intervals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const DefaultServerURL = "https://intervals.icu"

// Client represents the Intervals.icu API client
type Client struct {
	AthleteID  string
	APIKey     string
	ServerUrl  string
	httpClient *http.Client
}

// NewClient creates a new client with the default server URL
func NewClient(athleteID, apiKey string) *Client {
	return NewClientWithURL(athleteID, apiKey, DefaultServerURL)
}

// NewClientWithURL creates a new client with a custom server URL
func NewClientWithURL(athleteID, apiKey, serverURL string) *Client {
	return &Client{
		AthleteID:  athleteID,
		APIKey:     apiKey,
		ServerUrl:  serverURL,
		httpClient: &http.Client{},
	}
}

// makeRequest effectue une requête HTTP avec authentification
func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("erreur lors du marshal JSON: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.ServerUrl + endpoint
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de la requête: %w", err)
	}

	// Authentification basique avec AthleteID et APIKey
	req.SetBasicAuth("API_KEY", c.APIKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return c.httpClient.Do(req)
}
