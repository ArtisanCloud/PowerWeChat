package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type BaseInfo struct {
	DeviceType string `json:"device_type"`
	DeviceID   string `json:"device_id"`
}

type Resp struct {
	BaseInfo *BaseInfo `json:"base_info"`
	ErrCode  int       `json:"errcode"`
	ErrMsg   string    `json:"errmsg"`
}

type ResponseDeviceAuthorize struct {
	response.ResponseOfficialAccount

	Resp []*Resp `json:"resp"`
}
