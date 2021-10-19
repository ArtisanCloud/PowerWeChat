package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	*response.ResponseWork

	OpenID string `json:"openid"`
}
