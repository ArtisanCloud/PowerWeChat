package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	*response.ResponseWork

	LivingInfo   *power.HashMap   `json:"living_info"`

}
