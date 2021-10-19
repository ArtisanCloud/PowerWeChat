package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAuthCheckEncryptedData struct {
	*response.ResponseMiniProgram

	Valid      bool    `json:"vaild"`
	CreateTime float64 `json:"create_time"`
}
