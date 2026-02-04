package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	intervals "github.com/tgrangeray/go.intervals"
)

// LoadConfigFromFile lit le fichier config.json et extrait athlete_id et api_key
func loadConfigFromFile(filename string) (athleteID string, apiKey string, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", "", fmt.Errorf("erreur lors de la lecture du fichier %s: %w", filename, err)
	}

	var configData struct {
		AthleteID string `json:"athlete_id"`
		APIKey    string `json:"api_key"`
	}

	err = json.Unmarshal(data, &configData)
	if err != nil {
		return "", "", fmt.Errorf("erreur lors du parsing JSON: %w", err)
	}

	return configData.AthleteID, configData.APIKey, nil
}

func main() {
	// Charger la configuration depuis le fichier JSON
	athleteID, apiKey, err := loadConfigFromFile("../../config.json")
	if err != nil {
		log.Fatal("Erreur lors du chargement de la configuration:", err)
	}

	// Créer un nouveau client avec les valeurs extraites
	client := intervals.NewClient(athleteID, apiKey)

	fmt.Printf("Client créé avec succès:\n")
	fmt.Printf("AthleteID: %s\n", client.AthleteID)
	fmt.Printf("APIKey: %s\n", client.APIKey)
	fmt.Printf("ServerUrl: %s\n", client.ServerUrl)

	// Récupérer et afficher les informations de l'athlète
	athlete, err := client.GetAthlete()
	if err != nil {
		log.Fatal("Erreur lors de la récupération de l'athlète:", err)
	}

	fmt.Printf("\n=== Informations de l'athlète ===\n")
	fmt.Printf("ID: %s\n", athlete.ID)
	fmt.Printf("Nom: %s\n", athlete.Name)
	fmt.Printf("Email: %s\n", athlete.Email)
	fmt.Printf("Fuseau horaire: %s\n", athlete.TimeZone)
	fmt.Printf("Pays: %s\n", athlete.Country)
	fmt.Printf("Poids: %.1f %s\n", athlete.Weight, athlete.WeightUnit)

	// Test de récupération des activités (dernières 10)
	fmt.Printf("\n=== Récupération des activités ===\n")
	activities, err := client.GetActivities("", "", 10)
	if err != nil {
		log.Printf("Erreur lors de la récupération des activités: %v", err)
	} else {
		fmt.Printf("Nombre d'activités récupérées: %d\n", len(activities))
		for i, activity := range activities {
			fmt.Printf("%d. %s (%s) - %s - Distance: %.2f km\n",
				i+1, activity.Name, activity.Type, activity.StartDateLocal, activity.Distance/1000)
		}
	}
}