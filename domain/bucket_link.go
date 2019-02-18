package domain

const KeyPrefixLink = "link"

type Link struct {
	Bucket

	Value LinkValue `json:"value"`
}

type LinkValue struct {
	Structure string `json:"structure"`
}
