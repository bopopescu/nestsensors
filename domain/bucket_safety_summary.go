package domain

const KeyPrefixSafetySummary = "safety_summary"

type SafetySummary struct {
	Bucket

	Value interface{} `json:"value"`
}
