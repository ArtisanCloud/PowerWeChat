package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetPerformance struct {
	response.ResponseMiniProgram
	DefaultTimeData string `json:"default_time_data"`
	CompareTimeData string `json:"compare_time_data"`
}
