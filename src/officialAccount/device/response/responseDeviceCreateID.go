package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type RespMsg struct {
	RetCode   int    `json:"ret_code"`
	ErrorInfo string `json:" error_info"`
}

type ResponseDeviceCreateID struct {
	response.ResponseOfficialAccount

	RespMsg *RespMsg `json:"resp_msg"`
}
