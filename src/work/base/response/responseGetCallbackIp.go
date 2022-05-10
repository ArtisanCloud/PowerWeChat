package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGetCallBackIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWork
}
