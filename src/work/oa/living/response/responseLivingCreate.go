package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLivingCreate struct {
	*response.ResponseWork

	LivingID int `json:"livingid"`
}
