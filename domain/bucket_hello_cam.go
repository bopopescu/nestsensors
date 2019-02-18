package domain

const KeyPrefixHelloCam = "quartz"

type HelloCam struct {
	Bucket

	Value HelloCamValue `json:"value"`
}

type HelloCamValue struct {
	ActivationTime          int64       `json:"activation_time"`
	AudioInputEnabled       bool        `json:"audio_input_enabled"`
	CameraType              int         `json:"camera_type"`
	Capabilities            []string    `json:"capabilities"`
	CvrEnrolled             string      `json:"cvr_enrolled"`
	Description             string      `json:"description"`
	DirectNexustalkHost     string      `json:"direct_nexustalk_host"`
	DownloadHost            string      `json:"download_host"`
	FabricID                string      `json:"fabric_id"`
	IPAddress               string      `json:"ip_address"`
	LastConnectTime         int64       `json:"last_connect_time"`
	LastDisconnectReason    string      `json:"last_disconnect_reason"`
	LastDisconnectTime      int64       `json:"last_disconnect_time"`
	LiveStreamHost          string      `json:"live_stream_host"`
	MacAddress              string      `json:"mac_address"`
	Model                   string      `json:"model"`
	NexusAPIHTTPServerURL   string      `json:"nexus_api_http_server_url"`
	PreviewStreamingEnabled bool        `json:"preview_streaming_enabled"`
	PublicShareEnabled      bool        `json:"public_share_enabled"`
	RecordedStreamHost      string      `json:"recorded_stream_host"`
	SerialNumber            string      `json:"serial_number"`
	SnapshotURL             SnapshotURL `json:"snapshot_url"`
	SoftwareVersion         string      `json:"software_version"`
	StreamingState          string      `json:"streaming_state"`
	StructureID             string      `json:"structure_id"`
	WeaveDeviceID           string      `json:"weave_device_id"`
	WebsocketNexustalkHost  string      `json:"websocket_nexustalk_host"`
	WhereID                 string      `json:"where_id"`
	WwnStreamHost           string      `json:"wwn_stream_host"`
}

type SnapshotURL struct {
	SnapshotURLPrefix string `json:"snapshot_url_prefix"`
	SnapshotURLSuffix string `json:"snapshot_url_suffix"`
}
