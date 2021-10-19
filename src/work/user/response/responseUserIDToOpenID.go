package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	*response.ResponseWork

	OpenID string `json:"openid"`
}
