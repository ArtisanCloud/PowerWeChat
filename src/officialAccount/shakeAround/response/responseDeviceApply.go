package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type DeviceApplyData struct {
	ApplyID      int    `json:"apply_id"`
	AuditStatus  int    `json:"audit_status"`
	AuditComment string `json:"audit_comment"`
}

type ResponseDeviceApply struct {
	* response.ResponseOfficialAccount

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
	* response.ResponseOfficialAccount

	Data *DeviceApplyDataStatus `json:"data"`

}


// -------------------------------------------------------





// -------------------------------------------------------





// -------------------------------------------------------



