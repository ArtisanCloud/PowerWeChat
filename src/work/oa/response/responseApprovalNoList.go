package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseApprovalNoList struct {
	response.ResponseWork
	SpNoList      []string `json:"sp_no_list"`
	NewNextCursor string   `json:"new_next_cursor"`
}
