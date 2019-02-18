package domain

// Bucket serves as a generic base type for bucket-derived types.
type Bucket struct {
	ObjectKey       string `json:"object_key"`
	ObjectRevision  int    `json:"object_revision"`
	ObjectTimestamp int64  `json:"object_timestamp"`
}

// keys maps object key prefixes to device, setting, and information types.
var keys = map[string]interface{}{
	KeyPrefixBuckets:                 Buckets{},
	KeyPrefixDemandResponse:          DemandResponse{},
	KeyPrefixDetectInfo:              DetectInfo{},
	KeyPrefixDetectLocales:           DetectLocales{},
	KeyPrefixDetectSelfTest:          DetectSelfTest{},
	KeyPrefixDetectState:             DetectState{},
	KeyPrefixDeviceAlertDialog:       DeviceAlertDialog{},
	KeyPrefixGeofenceInfo:            GeofenceInfo{},
	KeyPrefixHelloCam:                HelloCam{},
	KeyPrefixLearningThermostat:      LearningThermostat{},
	KeyPrefixLearningThermostatState: LearningThermostatState{},
	KeyPrefixLink:                    Link{},
	KeyPrefixMessageCenter:           MessageCenter{},
	KeyPrefixMessage:                 Message{},
	KeyPrefixMetadata:                Metadata{},
	KeyPrefixOccupancy:               Occupancy{},
	KeyPrefixRCSSettings:             RCSSettings{},
	KeyPrefixSafetySummary:           SafetySummary{},
	KeyPrefixThermostatActivity:      ThermostatActivity{},
	KeyPrefixStructureHistory:        StructureHistory{},
	KeyPrefixStructureMetadata:       StructureMetadata{},
	KeyPrefixStructure:               Structure{},
	KeyPrefixTemperatureSensor:       TemperatureSensor{},
	KeyPrefixTrack:                   Track{},
	KeyPrefixTrip:                    Trip{},
	KeyPrefixTuneups:                 Tuneups{},
	KeyPrefixUserAlertDialog:         UserAlertDialog{},
	KeyPrefixUserSettings:            UserSettings{},
	KeyPrefixUser:                    User{},
	KeyPrefixWhere:                   Where{},
	KeyPrefixWidgetTrack:             WidgetTrack{},
}
