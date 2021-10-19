package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCode2Session struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`

	*response.ResponseMiniProgram
}
