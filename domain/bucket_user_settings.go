package domain

const KeyPrefixUserSettings = "user_settings"

type UserSettings struct {
	Bucket

	Value UserSettingsValue `json:"value"`
}

type UserSettingsValue struct {
	AppSwuMinimumVersion struct {
		Pairing int `json:"pairing"`
	} `json:"app_swu_minimum_version"`
	EmailVerified                 bool   `json:"email_verified"`
	HasLocationHistory            bool   `json:"has_location_history"`
	Lang                          string `json:"lang"`
	LocationPrimaryDevice         string `json:"location_primary_device"`
	MaxKryptonites                int    `json:"max_kryptonites"`
	MaxKryptonitesPerStructure    int    `json:"max_kryptonites_per_structure"`
	MaxKryptonitesPerThermostat   int    `json:"max_kryptonites_per_thermostat"`
	MaxSmokeDetectors             int    `json:"max_smoke_detectors"`
	MaxSmokeDetectorsPerStructure int    `json:"max_smoke_detectors_per_structure"`
	MaxStructures                 int    `json:"max_structures"`
	MaxThermostats                int    `json:"max_thermostats"`
	MaxThermostatsPerStructure    int    `json:"max_thermostats_per_structure"`
	MaxWwnDevices                 int    `json:"max_wwn_devices"`
	MaxWwnDevicesPerStructure     int    `json:"max_wwn_devices_per_structure"`
	ReceiveMarketingEmails        bool   `json:"receive_marketing_emails"`
	ReceiveNestEmails             bool   `json:"receive_nest_emails"`
	ReceiveSupportEmails          bool   `json:"receive_support_emails"`
	State2Fa                      string `json:"state_2fa"`
	TosAcceptedVersion            int64  `json:"tos_accepted_version"`
	TosCurrentVersion             int64  `json:"tos_current_version"`
	TosMinimumVersion             int64  `json:"tos_minimum_version"`
}
