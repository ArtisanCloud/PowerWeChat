package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCheckEncryptedMsg struct {
	Valid      bool    `json:"vaild"`
	CreateTime float64 `json:"create_time"`

	*response.ResponseMiniProgram
}
