# go-intervals

Client Go pour l'API Intervals.icu v1.0.0

Intervals Icu API specifications : https://intervals.icu/api-docs.html

## Installation

```bash
go get github.com/tgrangeray/go-intervals
```

## Usage

### Utilisation comme bibliothèque

```go
package main

import (
    "fmt"
    "log"
    
    intervals "github.com/tgrangeray/go-intervals"
)

func main() {
    // Créer un nouveau client
    client := intervals.NewClient("your_athlete_id", "your_api_key")
    
    // Récupérer les informations de l'athlète
    athlete, err := client.GetAthlete()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Athlète: %s\n", athlete.Name)
    
    // Récupérer les activités
    activities, err := client.GetActivities(map[string]string{
        "limit": "10",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Nombre d'activités: %d\n", len(activities))
}
```

### Utilisation comme programme autonome

Pour tester le client directement :

```bash
cd cmd/intervals
go run main.go
```

### Exemple d'utilisation

Consultez le fichier [example/main.go](example/main.go) pour un exemple complet d'utilisation.

### Id d'athlète et clés d'API

Le client utilise l'athlete_id et l'api_key d'Intervals.icu. Vous trouverez ces informations au bas de la page des Paramètres.


### Tests

Les test utilisent un fichier `config.json` avec vos identifiants :

```json
{
  "athlete_id": "votre_athlete_id",
  "api_key": "votre_api_key"
}
```

### Exemple d'usage

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // Création du client depuis config.json
    client := NewClient(athlete_id, api_key)

    // Récupération des activités des 30 derniers jours
    activities, err := client.GetActivities("2024-01-01", "", 50)
    if err != nil {
        log.Fatal("Erreur:", err)
    }
    
    fmt.Printf("Nombre d'activités: %d\n", len(activities))
}
```

## Méthodes disponibles

| Méthode | Description |
|---------|-------------|
| `GetAthlete()` | Get the athlete with sportSettings and custom_items |
| `GetActivities(oldest, newest, limit)` | List activities for a date range in desc date order |
| `GetActivity(activityID)` | Get an activity |
| `GetEvents(oldest, newest)` | Get events (planned workouts, notes etc.) from athlete's calendar |
| `CreateEvent(event)` | Create an event (planned workout, note etc.) on the athlete's calendar |
| `UpdateEvent(eventID, event)` | Update an event (planned workout, note etc.) |
| `DeleteEvent(eventID)` | Delete an event (planned workout, notes etc.) from an athlete's calendar |
| `GetWellness(date)` | Get wellness record for date (local ISO-8601 day) |
| `UpdateWellness(date, wellness)` | Update the wellness record for the date (ISO-8601) |
| `GetWorkouts()` | List all the workouts in the athlete's library |
| `GetWorkout(workoutID)` | Get a workout |
| `CreateWorkout(workout)` | Create a new workout in a folder or plan in the athlete's workout library |
| `GetSportSettings()` | List sport settings for the athlete |
| `GetPowerCurves(from, to)` | List best power curves for the athlete |

## Structures de données

Le client utilise des structures Go typées pour toutes les réponses API :

- `Athlete` : Informations complètes de l'athlète
- `Activity` : Données d'activité avec métriques détaillées  
- `Event` : Événements du calendrier (workouts planifiés, notes)
- `Wellness` : Données de santé quotidiennes (FC repos, HRV, poids, etc.)
- `Workout` : Structure des workouts avec étapes
- `SportSettings` : Paramètres par sport (FTP, zones, etc.)
- `PowerCurves` : Courbes de puissance avec données historiques
