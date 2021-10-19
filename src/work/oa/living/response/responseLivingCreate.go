package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseLivingCreate struct {
	*response.ResponseWork

	LivingID int `json:"livingid"`
}
