package domain

const KeyPrefixDetectLocales = "topaz_resource"

type DetectLocales struct {
	Bucket

	Value DetectLocalesValue `json:"value"`
}

type DetectLocalesValue struct {
	SupportedLocales []string `json:"supported_locales"`
	Wheres           []string `json:"wheres"`
}
