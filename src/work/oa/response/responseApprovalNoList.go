package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseApprovalNoList struct {
	*response.ResponseWork
	SpNoList []string `json:"sp_no_list"`
}
