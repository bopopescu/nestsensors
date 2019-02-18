package domain

const KeyPrefixTrip = "trip"

type Trip struct {
	Bucket

	Value TripValue `json:"value"`
}

type TripValue struct {
	Trips []interface{} `json:"trips"`
}
