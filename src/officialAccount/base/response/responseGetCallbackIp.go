package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGetCallBackIP struct {
	*response.ResponseOfficialAccount

	IPList []string `json:"ip_list"`
}
