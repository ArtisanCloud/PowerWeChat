package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetPerformanceData struct {
	Data *power.HashMap `json:"data"`

	*response.ResponseMiniProgram
}
