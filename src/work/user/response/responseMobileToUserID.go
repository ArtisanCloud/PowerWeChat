package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMobileToUserID struct {
	response.ResponseWork
	UserID string `json:"userid"`
}
