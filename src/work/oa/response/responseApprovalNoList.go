package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseApprovalNoList struct {
	*response.ResponseWork
	SpNoList []string `json:"sp_no_list"`
}
