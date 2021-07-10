package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	response.ResponseWX
	OpenID string `json:"openid"`
}
