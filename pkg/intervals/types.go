package intervals

// Response structures based on OpenAPI spec

type Athlete struct {
	ID               string                   `json:"id"`
	Name             string                   `json:"name"`
	Email            string                   `json:"email"`
	ProfileImageURL  string                   `json:"profile_image_url"`
	TimeZone         string                   `json:"time_zone"`
	Birthday         string                   `json:"birthday"`
	Gender           string                   `json:"gender"`
	Country          string                   `json:"country"`
	Weight           float64                  `json:"weight"`
	WeightUnit       string                   `json:"weight_unit"`
	FTPHistory       []map[string]interface{} `json:"ftp_history"`
	MaxHRHistory     []map[string]interface{} `json:"max_hr_history"`
	RestingHRHistory []map[string]interface{} `json:"resting_hr_history"`
	LTHRHistory      []map[string]interface{} `json:"lthr_history"`
	CustomItems      []CustomItem             `json:"custom_items"`
	SportSettings    []SportSettings          `json:"sport_settings"`
}

type Activity struct {
	ID                   int64     `json:"id"`
	StartDateLocal       string    `json:"start_date_local"`
	Type                 string    `json:"type"`
	IcuAthleteID         string    `json:"icu_athlete_id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Distance             float64   `json:"distance"`
	MovingTime           int       `json:"moving_time"`
	ElapsedTime          int       `json:"elapsed_time"`
	TotalElevationGain   float64   `json:"total_elevation_gain"`
	Trainer              bool      `json:"trainer"`
	Commute              bool      `json:"commute"`
	MaxSpeed             float64   `json:"max_speed"`
	AverageSpeed         float64   `json:"average_speed"`
	HasHeartRate         bool      `json:"has_heartrate"`
	MaxHeartRate         int       `json:"max_heartrate"`
	AverageHeartRate     int       `json:"average_heartrate"`
	HasPower             bool      `json:"has_power"`
	MaxPower             int       `json:"max_power"`
	AveragePower         int       `json:"average_power"`
	WeightedAvgPower     int       `json:"weighted_avg_power"`
	NormalizedPower      int       `json:"normalized_power"`
	TrainingLoad         float64   `json:"training_load"`
	IntensityFactor      float64   `json:"intensity_factor"`
	TSS                  float64   `json:"tss"`
	WorkoutCode          string    `json:"workout_code"`
	Tags                 []string  `json:"tags"`
	Feel                 int       `json:"feel"`
	PerceivedExertion    int       `json:"perceived_exertion"`
	TrimmedStartDate     string    `json:"trimmed_start_date"`
	TrimmedEndDate       string    `json:"trimmed_end_date"`
	AutoCalcPower        bool      `json:"auto_calc_power"`
	Paired               bool      `json:"paired"`
	PairedEventID        int64     `json:"paired_event_id"`
	ExternalID           string    `json:"external_id"`
	IcuIgnoreTime        bool      `json:"icu_ignore_time"`
	IcuIgnoreHR          bool      `json:"icu_ignore_hr"`
	IcuIgnorePower       bool      `json:"icu_ignore_power"`
	IcuTrainingLoad      float64   `json:"icu_training_load"`
	IcuHrLoad            float64   `json:"icu_hr_load"`
	IcuPowerLoad         float64   `json:"icu_power_load"`
	IcuDistance          float64   `json:"icu_distance"`
	IcuMovingTime        int       `json:"icu_moving_time"`
	IcuJoules            float64   `json:"icu_joules"`
	IcuDJoules           float64   `json:"icu_d_joules"`
	Gap                  float64   `json:"gap"`
	GradientAdjustedPace float64   `json:"gradient_adjusted_pace"`
	IcuFtp               int       `json:"icu_ftp"`
	IcuWKg               float64   `json:"icu_w_kg"`
	IcuHrZones           []int     `json:"icu_hr_zones"`
	IcuPowerZones        []int     `json:"icu_power_zones"`
	IcuSwimPaceZones     []float64 `json:"icu_swim_pace_zones"`
}

type Event struct {
	ID                int64    `json:"id"`
	StartDateLocal    string   `json:"start_date_local"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Type              string   `json:"type"`
	Category          string   `json:"category"`
	Tags              []string `json:"tags"`
	IcuAthleteID      string   `json:"icu_athlete_id"`
	ShowAsNote        bool     `json:"show_as_note"`
	Color             string   `json:"color"`
	ScheduledLoad     float64  `json:"scheduled_load"`
	CompletedLoad     float64  `json:"completed_load"`
	Done              bool     `json:"done"`
	ExternalID        string   `json:"external_id"`
	WorkoutDoc        string   `json:"workout_doc"`
	MovingTime        int      `json:"moving_time"`
	Distance          float64  `json:"distance"`
	TrainingLoad      float64  `json:"training_load"`
	Feel              int      `json:"feel"`
	PerceivedExertion int      `json:"perceived_exertion"`
	IcuTrainingLoad   float64  `json:"icu_training_load"`
	IcuHrLoad         float64  `json:"icu_hr_load"`
	IcuPowerLoad      float64  `json:"icu_power_load"`
	PairedActivityID  int64    `json:"paired_activity_id"`
	CreatedWithApp    string   `json:"created_with_app"`
	HideFromAthlete   bool     `json:"hide_from_athlete"`
	AthleteCannotEdit bool     `json:"athlete_cannot_edit"`
}

type Wellness struct {
	ID                   string                 `json:"id"`
	IcuAthleteID         string                 `json:"icu_athlete_id"`
	RestingHR            int                    `json:"resting_hr"`
	HRV                  float64                `json:"hrv"`
	Weight               float64                `json:"weight"`
	SleepSecs            int                    `json:"sleep_secs"`
	Soreness             int                    `json:"soreness"`
	Fatigue              int                    `json:"fatigue"`
	Stress               int                    `json:"stress"`
	Mood                 int                    `json:"mood"`
	Motivation           int                    `json:"motivation"`
	Injury               int                    `json:"injury"`
	SpO2                 float64                `json:"spo2"`
	SystolicBP           int                    `json:"systolic_bp"`
	DiastolicBP          int                    `json:"diastolic_bp"`
	Hydration            int                    `json:"hydration"`
	OverallFeel          int                    `json:"overall_feel"`
	SleepQuality         int                    `json:"sleep_quality"`
	MenstrualPhase       string                 `json:"menstrual_phase"`
	MenstrualFlow        int                    `json:"menstrual_flow"`
	Contraceptive        bool                   `json:"contraceptive"`
	Supplements          string                 `json:"supplements"`
	InjuryType           string                 `json:"injury_type"`
	InjuryLocation       string                 `json:"injury_location"`
	Comments             string                 `json:"comments"`
	Updated              string                 `json:"updated"`
	CustomWellnessFields map[string]interface{} `json:"custom_wellness_fields"`
}

type Workout struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Sport        string        `json:"sport"`
	Steps        []WorkoutStep `json:"steps"`
	Tags         []string      `json:"tags"`
	IcuAthleteID string        `json:"icu_athlete_id"`
	ExternalID   string        `json:"external_id"`
	CreatedDate  string        `json:"created_date"`
	UpdatedDate  string        `json:"updated_date"`
	LoadTarget   float64       `json:"load_target"`
	MovingTime   int           `json:"moving_time"`
	TotalTime    int           `json:"total_time"`
	Distance     float64       `json:"distance"`
}

type WorkoutStep struct {
	Duration     int           `json:"duration"`
	DurationType string        `json:"duration_type"`
	TargetType   string        `json:"target_type"`
	TargetLow    float64       `json:"target_low"`
	TargetHigh   float64       `json:"target_high"`
	Zone         int           `json:"zone"`
	Text         string        `json:"text"`
	Steps        []WorkoutStep `json:"steps"`
}

type SportSettings struct {
	ID            int64                  `json:"id"`
	IcuAthleteID  string                 `json:"icu_athlete_id"`
	Sport         string                 `json:"sport"`
	FTP           int                    `json:"ftp"`
	LTHR          int                    `json:"lthr"`
	MaxHR         int                    `json:"max_hr"`
	RestingHR     int                    `json:"resting_hr"`
	Weight        float64                `json:"weight"`
	PowerZones    []int                  `json:"power_zones"`
	HRZones       []int                  `json:"hr_zones"`
	PaceZones     []float64              `json:"pace_zones"`
	SwimPaceZones []float64              `json:"swim_pace_zones"`
	BikeWeight    float64                `json:"bike_weight"`
	WheelSize     int                    `json:"wheel_size"`
	Cda           float64                `json:"cda"`
	Crr           float64                `json:"crr"`
	PowerModel    map[string]interface{} `json:"power_model"`
}

type CustomItem struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Units        string `json:"units"`
	DisplayIndex int    `json:"display_index"`
	ShowChart    bool   `json:"show_chart"`
	ChartType    string `json:"chart_type"`
	ChartColor   string `json:"chart_color"`
}

type PowerCurveData struct {
	Secs  []int     `json:"secs"`
	Watts []float64 `json:"watts"`
	Dates []string  `json:"dates"`
}

type PowerCurves struct {
	Curve    PowerCurveData `json:"curve"`
	AllTime  PowerCurveData `json:"all_time"`
	Recent   PowerCurveData `json:"recent"`
	LastYear PowerCurveData `json:"last_year"`
}
