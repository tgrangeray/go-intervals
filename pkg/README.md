# Package Structure

This package has been restructured to follow Go best practices with a modular design.

## Structure

```
pkg/
└── intervals/
    ├── client.go           # Client initialization and core HTTP handling
    ├── types.go           # All type definitions and structs
    ├── athlete.go         # Athlete-related endpoints
    ├── activities.go      # Activity-related endpoints
    ├── events.go          # Event/calendar endpoints
    ├── wellness.go        # Wellness data endpoints
    ├── workouts.go        # Workout endpoints
    ├── sport_settings.go  # Sport settings endpoints
    └── power_curves.go    # Power curve endpoints
```

## Usage

### New recommended approach (modular package):

```go
import "github.com/tgrangeray/go.intervals/pkg/intervals"

client := intervals.NewClient(athleteID, apiKey)
athlete, err := client.GetAthlete()
```

### Backwards compatible approach (legacy):

```go
import "github.com/tgrangeray/go.intervals"

client := intervals.NewClient(athleteID, apiKey)
athlete, err := client.GetAthlete()
```

## Benefits of the New Structure

1. **Modularity**: Each functionality is separated into its own file
2. **Maintainability**: Easier to locate and modify specific features
3. **Testing**: Each module can be tested independently
4. **Documentation**: Clear separation of concerns
5. **Standards**: Follows Go project layout conventions with `pkg/` directory

## Migration Guide

No changes are required for existing code. The root `intervals.go` file provides full backwards compatibility by re-exporting all types and functions from the new package structure.

For new projects, we recommend using the modular import path: `github.com/tgrangeray/go.intervals/pkg/intervals`