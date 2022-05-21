package request

type RequestDeviceApply struct {
	Quantity    int    `json:"quantity"`
	ApplyReason string `json:"apply_reason"`
	Comment     string `json:"comment"`
	PoiID       int    `json:"poi_id"`
}


// ---------------------------------------------------

type RequestDeviceApplyStatus struct {
	ApplyID int `json:"apply_id"`
}

// ---------------------------------------------------

type DeviceIdentifier struct {
	DeviceID int    `json:"device_id"`
	UUID     string `json:"uuid"`
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
}

type RequestDeviceUpdate struct {
	DeviceIdentifier *DeviceIdentifier  `json:"device_identifier"`
	Comment string `json:"comment"`
}