package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGetCallBackIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWork
}
