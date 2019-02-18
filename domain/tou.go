package domain

const KeyPrefixTOU = "tou"

type TOU struct {
	Bucket

	Value interface{} `json:"value"`
}
