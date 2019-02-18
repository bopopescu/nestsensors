package domain

type Auth struct {
	AccessToken  string `json:"access_token"`
	Email        string `json:"email"`
	ExpiresIn    string `json:"expires_in"`
	TwoFAEnabled bool   `json:"2fa_enabled"`
	IsSuperuser  bool   `json:"is_superuser"`
	Language     string `json:"language"`
	User         string `json:"user"`
	IsStaff      bool   `json:"is_staff"`
	URLs         URLs   `json:"urls"`
	Limits       Limits `json:"limits"`
	Weave        Weave  `json:"weave"`
	UserID       string `json:"userid"`
}
