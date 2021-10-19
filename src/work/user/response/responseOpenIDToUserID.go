package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
