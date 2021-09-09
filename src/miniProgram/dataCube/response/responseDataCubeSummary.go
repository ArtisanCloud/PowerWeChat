package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	*response.ResponseMiniProgram

	List     []*power.HashMap `json:"list"`
}
