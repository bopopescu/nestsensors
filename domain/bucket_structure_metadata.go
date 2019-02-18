package domain

const KeyPrefixStructureMetadata = "structure_metadata"

type StructureMetadata struct {
	Bucket

	Value StructureMetadataValue `json:"value"`
}

type StructureMetadataValue struct {
	EntitlementsAPIPath     string `json:"entitlements_api_path"`
	StructureHistoryAPIPath string `json:"structure_history_api_path"`
}
