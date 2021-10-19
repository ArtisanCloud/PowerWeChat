package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGetCallBackIp struct {
	*response.ResponseOfficialAccount
	IPList []string `json:"ip_list"`
}
