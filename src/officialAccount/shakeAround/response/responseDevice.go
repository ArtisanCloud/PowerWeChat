package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type DeviceApplyData struct {
	ApplyID      int    `json:"apply_id"`
	AuditStatus  int    `json:"audit_status"`
	AuditComment string `json:"audit_comment"`
}

type ResponseDeviceApply struct {
	*response.ResponseOfficialAccount

	Data *DeviceApplyData `json:"data"`
}

// -------------------------------------------------------

type DeviceApplyDataStatus struct {
	ApplyTime    int    `json:"apply_time"`
	AuditComment string `json:"audit_comment"`
	AuditStatus  int    `json:"audit_status"`
	AuditTime    int    `json:"audit_time"`
}

type ResponseDeviceApplyStatus struct {
	*response.ResponseOfficialAccount

	Data *DeviceApplyDataStatus `json:"data"`
}

// -------------------------------------------------------

type ResponseDeviceBindPoi struct {
	*response.ResponseOfficialAccount

	Data interface{} `json:"data"`
}

// -------------------------------------------------------

type Device struct {
	Comment        string `json:"comment"`
	DeviceID       int    `json:"device_id"`
	Major          int    `json:"major"`
	Minor          int    `json:"minor"`
	Status         int    `json:"status"`
	LastActiveTime int    `json:"last_active_time"`
	PoiID          int    `json:"poi_id"`
	UUID           string `json:"uuid"`
	PoiAppID       string `json:"poi_appid,omitempty"`
	ShakePv        int    `json:"shake_pv"`
	ShakeUv        int    `json:"shake_uv"`
	ClickPv        int    `json:"click_pv"`
	ClickUv        int    `json:"click_uv"`
}

type SearchData struct {
	Devices    []*Device `json:"devices"`
	TotalCount int       `json:"total_count"`
}

type ResponseDeviceSearch struct {
	*response.ResponseOfficialAccount

	Data *SearchData `json:"data"`
}

// -------------------------------------------------------
