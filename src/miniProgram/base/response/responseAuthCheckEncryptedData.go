package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAuthCheckEncryptedData struct {
	*response.ResponseMiniProgram

	Valid      bool    `json:"vaild"`
	CreateTime float64 `json:"create_time"`
}
