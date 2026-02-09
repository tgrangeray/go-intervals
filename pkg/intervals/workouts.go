package intervals

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetWorkouts récupère la liste des workouts
func (c *Client) GetWorkouts() ([]Workout, error) {
	resp, err := c.makeRequest("GET", "/api/v1/athlete/"+c.AthleteID+"/workouts", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result []Workout
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// GetWorkout récupère un workout spécifique
func (c *Client) GetWorkout(workoutID string) (*Workout, error) {
	resp, err := c.makeRequest("GET", "/api/v1/athlete/"+c.AthleteID+"/workouts/"+workoutID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Workout
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}

// CreateWorkout crée un nouveau workout
func (c *Client) CreateWorkout(workout Workout) (*Workout, error) {
	resp, err := c.makeRequest("POST", "/api/v1/athlete/"+c.AthleteID+"/workouts", workout)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("erreur API: %s", resp.Status)
	}

	var result Workout
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
