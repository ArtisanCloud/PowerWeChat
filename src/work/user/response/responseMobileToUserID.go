package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMobileToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
