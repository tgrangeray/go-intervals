package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetPowerCurves récupère les courbes de puissance
func (c *Client) GetPowerCurves(from, to string) (*PowerCurves, error) {
	params := url.Values{}
	if from != "" {
		params.Add("from", from)
	}
	if to != "" {
		params.Add("to", to)
	}

	endpoint := "/api/v1/athlete/" + c.AthleteID + "/power-curves"
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

	var result PowerCurves
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
