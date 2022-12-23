package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Interface struct {
	ApiName    string `json:"api_name"`
	ApiChName  string `json:"api_ch_name"`
	ApiDesc    string `json:"api_desc"`
	Status     int    `json:"status"`
	ApiLink    string `json:"api_link"`
	GroupName  string `json:"group_name"`
	AuditID    int    `json:"audit_id,omitempty"`
	FailReason string `json:"fail_reason,omitempty"`
}

type ResponseGet struct {
	response.ResponseOpenPlatform

	InterfaceList []*Interface `json:"interface_list"`
}

type ResponseSet struct {
	response.ResponseOpenPlatform

	AuditId int `json:"audit_id"`
}
