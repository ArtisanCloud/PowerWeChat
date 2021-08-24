package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	response.ResponseWork
	UserID string `json:"userid"`
}
