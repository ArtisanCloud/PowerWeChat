package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	response.ResponseMiniProgram

	List []*power.HashMap `json:"list"`
}
