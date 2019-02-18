package domain

const KeyPrefixDetectState = "topaz"

type DetectState struct {
	Bucket

	Value DetectStateValue `json:"value"`
}

type DetectStateValue struct {
	AutoAway                     bool     `json:"auto_away"`
	BatteryHealthState           int      `json:"battery_health_state"`
	BatteryLevel                 int      `json:"battery_level"`
	BuzzerTestResults            int      `json:"buzzer_test_results"`
	CapabilityLevel              int      `json:"capability_level"`
	CertificationBody            int      `json:"certification_body"`
	CoBlameDuration              int      `json:"co_blame_duration"`
	CoBlameThreshold             int      `json:"co_blame_threshold"`
	CoPreviousPeak               int      `json:"co_previous_peak"`
	CoSequenceNumber             int      `json:"co_sequence_number"`
	CoStatus                     int      `json:"co_status"`
	ComponentAlsTestPassed       bool     `json:"component_als_test_passed"`
	ComponentBuzzerTestPassed    bool     `json:"component_buzzer_test_passed"`
	ComponentCoTestPassed        bool     `json:"component_co_test_passed"`
	ComponentHeatTestPassed      bool     `json:"component_heat_test_passed"`
	ComponentHumTestPassed       bool     `json:"component_hum_test_passed"`
	ComponentLedTestPassed       bool     `json:"component_led_test_passed"`
	ComponentPirTestPassed       bool     `json:"component_pir_test_passed"`
	ComponentSmokeTestPassed     bool     `json:"component_smoke_test_passed"`
	ComponentSpeakerTestPassed   bool     `json:"component_speaker_test_passed"`
	ComponentTempTestPassed      bool     `json:"component_temp_test_passed"`
	ComponentUsTestPassed        bool     `json:"component_us_test_passed"`
	ComponentWifiTestPassed      bool     `json:"component_wifi_test_passed"`
	CreationTime                 int64    `json:"creation_time"`
	DeviceBornOnDateUtcSecs      int      `json:"device_born_on_date_utc_secs"`
	DeviceExternalColor          string   `json:"device_external_color"`
	DeviceLocale                 string   `json:"device_locale"`
	FabricID                     string   `json:"fabric_id"`
	FactoryLoadedLanguages       string   `json:"factory_loaded_languages"`
	GestureHushEnable            bool     `json:"gesture_hush_enable"`
	HeadsUpEnable                bool     `json:"heads_up_enable"`
	HomeAlarmLinkCapable         bool     `json:"home_alarm_link_capable"`
	HomeAwayInput                bool     `json:"home_away_input"`
	HushedState                  bool     `json:"hushed_state"`
	InstalledLocale              string   `json:"installed_locale"`
	KlSoftwareVersion            string   `json:"kl_software_version"`
	LatestManualTestCancelled    bool     `json:"latest_manual_test_cancelled"`
	LatestManualTestEndUtcSecs   int      `json:"latest_manual_test_end_utc_secs"`
	LatestManualTestStartUtcSecs int      `json:"latest_manual_test_start_utc_secs"`
	LinePowerPresent             bool     `json:"line_power_present"`
	Model                        string   `json:"model"`
	NightLightContinuous         bool     `json:"night_light_continuous"`
	NightLightEnable             bool     `json:"night_light_enable"`
	NtpGreenLedBrightness        int      `json:"ntp_green_led_brightness"`
	NtpGreenLedEnable            bool     `json:"ntp_green_led_enable"`
	ProductID                    int      `json:"product_id"`
	ReplaceByDateUtcSecs         int      `json:"replace_by_date_utc_secs"`
	ResourceID                   string   `json:"resource_id"`
	SerialNumber                 string   `json:"serial_number"`
	SmokeSequenceNumber          int      `json:"smoke_sequence_number"`
	SmokeStatus                  int      `json:"smoke_status"`
	SoftwareVersion              string   `json:"software_version"`
	SpeakerTestResults           int      `json:"speaker_test_results"`
	SpokenWhereID                string   `json:"spoken_where_id"`
	SteamDetectionEnable         bool     `json:"steam_detection_enable"`
	StructureID                  string   `json:"structure_id"`
	ThreadIPAddress              []string `json:"thread_ip_address"`
	ThreadMacAddress             string   `json:"thread_mac_address"`
	WhereID                      string   `json:"where_id"`
	WifiIPAddress                string   `json:"wifi_ip_address"`
	WifiMacAddress               string   `json:"wifi_mac_address"`
	WifiRegulatoryDomain         string   `json:"wifi_regulatory_domain"`
	WiredLedEnable               bool     `json:"wired_led_enable"`
	WiredOrBattery               int      `json:"wired_or_battery"` // 1 == battery, presumably 2 == wired.
}
