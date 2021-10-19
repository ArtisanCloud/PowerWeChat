package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	*response.ResponseMiniProgram
	GrayReleasePlan   *power.HashMap `json:"gray_release_plan"`

}
