package domain

const KeyPrefixThermostatActivity = "schedule"

type ThermostatActivity struct {
	Bucket

	Value ThermostatActivityValue `json:"value"`
}

type ThermostatActivityValue struct {
	Days         map[string]ThermostatActivityEntry `json:"days"`
	Name         string                             `json:"name"`
	ScheduleMode string                             `json:"schedule_mode"`
	Ver          int                                `json:"ver"`
}

type ThermostatActivityEntry struct {
	EntryType  string  `json:"entry_type"`
	Temp       float64 `json:"temp"`
	Time       int     `json:"time"`
	TouchedAt  int     `json:"touched_at"`
	TouchedBy  int     `json:"touched_by"`
	TouchedTZO int     `json:"touched_tzo"`
	Type       string  `json:"type"`
}
