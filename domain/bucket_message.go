package domain

const KeyPrefixMessage = "message"

type Message struct {
	Bucket

	Value interface{} `json:"value"`
}
