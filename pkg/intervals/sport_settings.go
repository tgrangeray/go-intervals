package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSportSettings récupère les paramètres sport
func (c *Client) GetSportSettings() ([]SportSettings, error) {
	resp, err := c.makeRequest("GET", "/api/v1/athlete/"+c.AthleteID+"/sport-settings", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result []SportSettings
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
