package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetCallBackIp struct {
	*response.ResponseOfficialAccount
	IPList []string `json:"ip_list"`
}
