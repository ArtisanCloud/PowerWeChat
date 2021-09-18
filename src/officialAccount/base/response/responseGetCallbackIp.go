package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetCallBackIp struct {
	*response.ResponseOfficialAccount
	IPList []string `json:"ip_list"`
}
