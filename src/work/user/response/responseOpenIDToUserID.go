package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	response.ResponseWX
	UserID string `json:"userid"`
}
