package response

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseGetCallBackIp struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWX
}

