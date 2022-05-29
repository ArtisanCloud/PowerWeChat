package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseApprovalNoList struct {
	*response.ResponseWork
	SpNoList []string `json:"sp_no_list"`
}
