package domain

const KeyPrefixDeviceAlertDialog = "device_alert_dialog"

type DeviceAlertDialog struct {
	Bucket

	Value DeviceAlertDialogValue `json:"value"`
}

type DeviceAlertDialogValue struct {
	DialogData string `json:"dialog_data"`
	DialogID   string `json:"dialog_id"`
}
