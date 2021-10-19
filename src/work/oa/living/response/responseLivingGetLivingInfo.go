package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	*response.ResponseWork

	LivingInfo   *power.HashMap   `json:"living_info"`

}
