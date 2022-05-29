package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	*response.ResponseMiniProgram
	GrayReleasePlan *power.HashMap `json:"gray_release_plan"`
}
