package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	*response.ResponseMiniProgram

	List     []*power.HashMap `json:"list"`
}
