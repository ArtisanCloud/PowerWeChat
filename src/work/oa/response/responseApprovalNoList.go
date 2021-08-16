package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseApprovalNoList struct {
	*response.ResponseWork
	SpNoList []string `json:"sp_no_list"`
}
