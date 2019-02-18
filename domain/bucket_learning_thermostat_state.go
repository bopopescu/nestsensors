package domain

const KeyPrefixLearningThermostatState = "shared"

type LearningThermostatState struct {
	Bucket

	Value LearningThermostatStateValue `json:"value"`
}

type LearningThermostatStateValue struct {
	AutoAway                 int       `json:"auto_away"`
	AutoAwayLearning         string    `json:"auto_away_learning"`
	CanCool                  bool      `json:"can_cool"`
	CanHeat                  bool      `json:"can_heat"`
	CompressorLockoutEnabled bool      `json:"compressor_lockout_enabled"`
	CompressorLockoutTimeout int       `json:"compressor_lockout_timeout"`
	CurrentTemperature       float64   `json:"current_temperature"`
	HVACACState              bool      `json:"hvac_ac_state"`
	HVACAltHeatState         bool      `json:"hvac_alt_heat_state"`
	HVACAltHeatX2State       bool      `json:"hvac_alt_heat_x2_state"`
	HVACAuxHeaterState       bool      `json:"hvac_aux_heater_state"`
	HVACCoolX2State          bool      `json:"hvac_cool_x2_state"`
	HVACCoolX3State          bool      `json:"hvac_cool_x3_state"`
	HVACEmerHeatState        bool      `json:"hvac_emer_heat_state"`
	HVACFanState             bool      `json:"hvac_fan_state"`
	HVACHeatX2State          bool      `json:"hvac_heat_x2_state"`
	HVACHeatX3State          bool      `json:"hvac_heat_x3_state"`
	HVACHeaterState          bool      `json:"hvac_heater_state"`
	Name                     string    `json:"name"`
	TargetChangePending      bool      `json:"target_change_pending"`
	TargetTemperature        float64   `json:"target_temperature"`
	TargetTemperatureHigh    int       `json:"target_temperature_high"`
	TargetTemperatureLow     int       `json:"target_temperature_low"`
	TargetTemperatureType    string    `json:"target_temperature_type"`
	TouchedBy                TouchedBy `json:"touched_by"`
}

type TouchedBy struct {
	By     int    `json:"touched_by"`
	ID     string `json:"touched_id"`
	UserID string `json:"touched_user_id"`
}
