package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseActiveMessageCreateActiveID struct {
	response.ResponseMiniProgram
	ActivityID     string  `json:"activity_id"`
	ExpirationTime float64 `json:"expiration_time"`
}
