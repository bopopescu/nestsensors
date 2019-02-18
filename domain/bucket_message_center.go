package domain

const KeyPrefixMessageCenter = "message_center"

type MessageCenter struct {
	Bucket

	Value MessageCenterValue `json:"value"`
}

type MessageCenterValue struct {
	Messages MCMessages `json:"messages"`
}

type MCMessage struct {
	Dismissed  bool     `json:"dismissed"`
	ID         string   `json:"id"`
	Key        string   `json:"key"`
	Parameters []string `json:"parameters"`
	Priority   int      `json:"priority"`
	Read       bool     `json:"read"`
	ThreadID   int      `json:"thread_id"`
	Timestamp  int      `json:"timestamp"`
}

type MCMessages []MCMessage
