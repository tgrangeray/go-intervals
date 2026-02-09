// Package intervals provides a backwards compatibility layer.
// For new code, use github.com/tgrangeray/go.intervals/pkg/intervals directly.
package intervals

import "github.com/tgrangeray/go.intervals/pkg/intervals"

// Client is an alias for the intervals.Client type
type Client = intervals.Client

// Response types
type (
	Athlete        = intervals.Athlete
	Activity       = intervals.Activity
	Event          = intervals.Event
	Wellness       = intervals.Wellness
	Workout        = intervals.Workout
	WorkoutStep    = intervals.WorkoutStep
	SportSettings  = intervals.SportSettings
	CustomItem     = intervals.CustomItem
	PowerCurveData = intervals.PowerCurveData
	PowerCurves    = intervals.PowerCurves
)

// Constants
const DefaultServerURL = intervals.DefaultServerURL

// Constructor functions
var (
	NewClient        = intervals.NewClient
	NewClientWithURL = intervals.NewClientWithURL
)
