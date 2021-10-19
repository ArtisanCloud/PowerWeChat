package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
