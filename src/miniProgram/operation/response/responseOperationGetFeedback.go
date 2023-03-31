package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetFeedback struct {
	response.ResponseMiniProgram
	List     []string `json:"list"`
	TotalNum int64    `json:"total_num"`
}
