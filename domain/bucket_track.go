package domain

const KeyPrefixTrack = "track"

type Track struct {
	Bucket

	Value TrackValue `json:"value"`
}

type TrackValue struct {
	LastConnection int64  `json:"last_connection"`
	LastIP         string `json:"last_ip"`
	Online         bool   `json:"online"`
}
