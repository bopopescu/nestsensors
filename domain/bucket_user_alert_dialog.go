package domain

const KeyPrefixUserAlertDialog = "user_alert_dialog"

type UserAlertDialog struct {
	Bucket

	Value UserAlertDialogValue `json:"value"`
}

type UserAlertDialogValue struct {
	DialogData string `json:"dialog_data"`
	DialogID   string `json:"dialog_id"`
}
