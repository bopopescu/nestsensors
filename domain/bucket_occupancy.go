package domain

const KeyPrefixOccupancy = "occupancy"

type Occupancy struct {
	Bucket

	Value OccupancyValue `json:"value"`
}

type OccupancyValue struct {
	ElementDurationSeconds int   `json:"element_duration_seconds"`
	Occupancy              []int `json:"occupancy"`
}
