package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationRealTimeLogSearch struct {
	*response.ResponseMiniProgram
	Data *power.HashMap   `json:"data"`
	List []*power.HashMap `json:"list"`
}
