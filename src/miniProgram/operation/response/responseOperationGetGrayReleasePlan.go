package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	*response.ResponseMiniProgram
	GrayReleasePlan   *power.HashMap `json:"gray_release_plan"`

}
