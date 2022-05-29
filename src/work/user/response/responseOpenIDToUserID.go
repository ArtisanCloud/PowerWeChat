package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
