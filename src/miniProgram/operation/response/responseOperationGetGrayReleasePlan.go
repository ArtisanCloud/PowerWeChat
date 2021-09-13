package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	*response.ResponseMiniProgram
	GrayReleasePlan   *power.HashMap `json:"gray_release_plan"`

}
