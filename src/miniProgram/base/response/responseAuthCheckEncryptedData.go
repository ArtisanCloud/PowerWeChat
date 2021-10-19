package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAuthCheckEncryptedData struct {
	*response.ResponseMiniProgram

	Valid      bool    `json:"vaild"`
	CreateTime float64 `json:"create_time"`
}
