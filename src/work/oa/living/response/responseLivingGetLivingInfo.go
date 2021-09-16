package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	*response.ResponseWork

	LivingInfo   *power.HashMap   `json:"living_info"`

}
