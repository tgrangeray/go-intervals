package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetWellness récupère les données de wellness pour une date
func (c *Client) GetWellness(date string) (*Wellness, error) {
	resp, err := c.makeRequest("GET", "/api/v1/athlete/"+c.AthleteID+"/wellness/"+date, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Wellness
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}

// UpdateWellness met à jour les données de wellness pour une date
func (c *Client) UpdateWellness(date string, wellness Wellness) (*Wellness, error) {
	resp, err := c.makeRequest("PUT", "/api/v1/athlete/"+c.AthleteID+"/wellness/"+date, wellness)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Wellness
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
