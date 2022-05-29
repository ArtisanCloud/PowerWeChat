package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	*response.ResponseWork

	OpenID string `json:"openid"`
}
