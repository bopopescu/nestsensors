package domain

const KeyPrefixBuckets = "buckets"

type Buckets struct {
	Bucket

	Value BucketsValue `json:"value"`
}

type BucketsValue struct {
	Buckets []string `json:"buckets"`
}
