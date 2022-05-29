package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	*response.ResponseWork

	LivingInfo *power.HashMap `json:"living_info"`
}
