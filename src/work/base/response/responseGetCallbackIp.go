package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGetCallBackIp struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWork
}
