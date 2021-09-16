package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLivingGetLivingCode struct {
	*response.ResponseWork

	LivingCode int `json:"living_code"`
}
