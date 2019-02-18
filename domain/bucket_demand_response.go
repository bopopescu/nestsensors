package domain

const KeyPrefixDemandResponse = "demand_response"

type DemandResponse struct {
	Bucket

	Value interface{} `json:"value"`
}
