package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetEvents récupère les événements du calendrier
func (c *Client) GetEvents(oldest, newest string) ([]Event, error) {
	params := url.Values{}
	if oldest != "" {
		params.Add("oldest", oldest)
	}
	if newest != "" {
		params.Add("newest", newest)
	}

	endpoint := "/api/v1/athlete/" + c.AthleteID + "/events"
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

	var result []Event
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// CreateEvent crée un nouvel événement
func (c *Client) CreateEvent(event Event) (*Event, error) {
	resp, err := c.makeRequest("POST", "/api/v1/athlete/"+c.AthleteID+"/events", event)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Event
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}

// UpdateEvent met à jour un événement existant
func (c *Client) UpdateEvent(eventID string, event Event) (*Event, error) {
	resp, err := c.makeRequest("PUT", "/api/v1/athlete/"+c.AthleteID+"/events/"+eventID, event)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Event
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}

// DeleteEvent supprime un événement
func (c *Client) DeleteEvent(eventID string) error {
	resp, err := c.makeRequest("DELETE", "/api/v1/athlete/"+c.AthleteID+"/events/"+eventID, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API: %s", resp.Status)
	}

	return nil
}
