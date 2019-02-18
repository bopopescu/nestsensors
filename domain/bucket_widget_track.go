package domain

const KeyPrefixWidgetTrack = "widget_track"

// WidgetTrack not yet clear on what this is, exactly.
type WidgetTrack struct {
	Bucket

	Value WidgetTrackValue `json:"value"`
}

type WidgetTrackValue struct {
	LastConnection int64  `json:"last_connection"`
	LastIP         string `json:"last_ip"`
	Online         bool   `json:"online"`
}
