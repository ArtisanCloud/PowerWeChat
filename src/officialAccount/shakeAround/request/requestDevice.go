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

type RequestDeviceIdentifier struct {
	DeviceID int    `json:"device_id"`
	UUID     string `json:"uuid"`
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
}

type RequestDeviceSearch struct {
	Type              int `json:"type"`
	DeviceIdentifiers []*RequestDeviceIdentifier `json:"device_identifiers"`
	ApplyID  int `json:"apply_id"`
	LastSeen int `json:"last_seen"`
	Count    int `json:"count"`
}