package domain

const KeyPrefixTuneups = "tuneups"

type Tuneups struct {
	Bucket

	Value interface{} `json:"value"`
}
