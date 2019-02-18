package domain

const KeyPrefixDetectSelfTest = "safety"

// DetectSelfTest represents a Nest Detect smoke alarm.
type DetectSelfTest struct {
	Bucket

	Value DetectSelfTestValue `json:"value"`
}

type DetectSelfTestValue struct {
	AudioSelfTestEnabled            bool     `json:"audio_self_test_enabled"`
	AudioSelfTestEndUtcOffsetSecs   int      `json:"audio_self_test_end_utc_offset_secs"`
	AudioSelfTestParticipants       []string `json:"audio_self_test_participants"`
	AudioSelfTestStartUtcOffsetSecs int      `json:"audio_self_test_start_utc_offset_secs"`
	LastManualSelfTestCancelled     bool     `json:"last_manual_self_test_cancelled"`
	LastManualSelfTestEndUtcSecs    int      `json:"last_manual_self_test_end_utc_secs"`
	LastManualSelfTestStartUtcSecs  int      `json:"last_manual_self_test_start_utc_secs"`
	ManualSelfTestInProgress        bool     `json:"manual_self_test_in_progress"`
}
