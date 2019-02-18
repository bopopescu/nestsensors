package domain

const KeyPrefixMetadata = "metadata"

type Metadata struct {
	Bucket

	Value MetadataValue `json:"value"`
}

type MetadataValue struct {
	LastConnection int64  `json:"last_connection"`
	LastIP         string `json:"last_ip"`
	Online         bool   `json:"online"`
}
