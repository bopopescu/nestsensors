package domain

const KeyPrefixGeofenceInfo = "geofence_info"

type GeofenceInfo struct {
	Bucket

	Value GeofenceInfoValue `json:"value"`
}

type GeofenceInfoValue struct {
	CombinedPresence CombinedPresence `json:"combined_presence"`
	DeviceEvents     DeviceEvent      `json:"device_events"`
	Fences           Fences           `json:"fences"`
}

type CombinedPresence struct {
	Presence                    int   `json:"presence"`
	PresenceEvaluationTimestamp int64 `json:"presence_evaluation_timestamp"`
	RawPresence                 int   `json:"raw_presence"`
}

type DeviceEvent struct {
	DeviceID              string            `json:"device_id"`
	Events                DeviceFenceEvents `json:"events"`
	IsAttachedToOccupancy bool              `json:"is_attached_to_occupancy"`
	IsPrimaryDevice       bool              `json:"is_primary_device"`
}

type DeviceEvents []DeviceEvent

type DeviceFenceEvent struct {
	FenceID               string `json:"fence_id"`
	State                 int    `json:"state"`
	TimeOfStateTransition int64  `json:"time_of_state_transition"`
	TimeStateReported     int64  `json:"time_state_reported"`
	UserID                string `json:"user_id"`
}

type DeviceFenceEvents []DeviceFenceEvent

type Fence struct {
	Direction []int   `json:"direction"`
	FenceID   string  `json:"fence_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    int     `json:"radius"`
}

type Fences []Fence
