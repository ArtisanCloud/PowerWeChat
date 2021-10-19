package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMobileToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
