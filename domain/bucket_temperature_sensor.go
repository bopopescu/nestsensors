package domain

const KeyPrefixTemperatureSensor = "kryptonite"

type TemperatureSensor struct {
	Bucket

	Value TemperatureSensorValue `json:"value"`
}

type TemperatureSensorValue struct {
	BatteryLevel       int     `json:"battery_level"`
	CurrentTemperature float64 `json:"current_temperature"`
	LastUpdatedAt      int     `json:"last_updated_at"`
	Model              string  `json:"model"`
	SerialNumber       string  `json:"serial_number"`
	StructureID        string  `json:"structure_id"`
	WhereID            string  `json:"where_id"`
}
