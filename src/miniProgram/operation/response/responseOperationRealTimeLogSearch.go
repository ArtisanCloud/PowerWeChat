package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseOperationRealTimeLogSearch struct {
	*response.ResponseMiniProgram
	Data *power.HashMap   `json:"data"`
	List []*power.HashMap `json:"list"`
}
