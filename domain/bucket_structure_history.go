package domain

const KeyPrefixStructureHistory = "structure_history"

type StructureHistory struct {
	Bucket

	Value StructureHistoryValue `json:"value"`
}

type StructureHistoryValue struct {
	Events        []interface{} `json:"events"`
	Products      []interface{} `json:"products"`
	SchemaVersion string        `json:"schema_version"`
}
