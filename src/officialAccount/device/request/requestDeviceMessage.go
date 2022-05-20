package request

type RequestDeviceMessage struct {
	DeviceType string `json:"device_type"`
	DeviceID   string `json:"device_id"`
	OpenID     string `json:"open_id"`
	Content    string `json:"content"`
}
