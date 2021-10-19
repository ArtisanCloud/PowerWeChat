package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseLivingCreate struct {
	*response.ResponseWork

	LivingID int `json:"livingid"`
}
