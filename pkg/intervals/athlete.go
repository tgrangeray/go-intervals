package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAthlete récupère les informations de l'athlète
func (c *Client) GetAthlete() (*Athlete, error) {
	resp, err := c.makeRequest("GET", "/api/v1/athlete/"+c.AthleteID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Athlete
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
