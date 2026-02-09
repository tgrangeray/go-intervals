package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetActivities récupère la liste des activités
func (c *Client) GetActivities(oldest, newest string, limit int) ([]Activity, error) {
	params := url.Values{}
	if oldest != "" {
		params.Add("oldest", oldest)
	}
	if newest != "" {
		params.Add("newest", newest)
	}
	if limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}

	endpoint := "/api/v1/athlete/" + c.AthleteID + "/activities"
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	resp, err := c.makeRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result []Activity
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// GetActivity récupère une activité spécifique
func (c *Client) GetActivity(activityID string) (*Activity, error) {
	resp, err := c.makeRequest("GET", "/api/v1/activity/"+activityID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Activity
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
