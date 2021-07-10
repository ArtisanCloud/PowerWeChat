package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMobileToUserID struct {
	response.ResponseWX
	UserID string `json:"userid"`
}
