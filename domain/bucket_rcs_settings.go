package domain

const KeyPrefixRCSSettings = "rcs_settings"

type RCSSettings struct {
	Bucket

	Value RCSSettingsValue `json:"value"`
}

type RCSSettingsValue struct {
	ActiveRCSSensors     []string       `json:"active_rcs_sensors"`
	AssociatedRCSSensors []string       `json:"associated_rcs_sensors"`
	MultiroomActive      bool           `json:"multiroom_active"`
	MultiroomTemperature float64        `json:"multiroom_temperature"`
	RCSControlSetting    string         `json:"rcs_control_setting"`
	SensorInsights       SensorInsights `json:"sensor_insights"`
	SensorSchedule       SensorSchedule `json:"sensor_schedule"`
	ThermostatAlert      string         `json:"thermostat_alert"`
}

type SensorInsights struct {
	Insights Insights `json:"insights"`
}

type Insight struct {
	CreatedAt       int    `json:"created_at"`
	RangeTestResult string `json:"range_test_result"`
	SensorID        string `json:"sensor_id"`
	Value           int    `json:"value"`
}

type Insights []Insight

type SensorSchedule struct {
	Intervals Intervals `json:"intervals"`
}

type Interval struct {
	EndDay          int      `json:"end_day"`
	EndTime         int      `json:"end_time"`
	MultiroomActive bool     `json:"multiroom_active"`
	Sensors         []string `json:"sensors"`
	StartDay        int      `json:"start_day"`
	StartTime       int      `json:"start_time"`
}

type Intervals []Interval
