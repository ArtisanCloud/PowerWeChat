package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseLivingGetLivingCode struct {
	*response.ResponseWork

	LivingCode int `json:"living_code"`
}
