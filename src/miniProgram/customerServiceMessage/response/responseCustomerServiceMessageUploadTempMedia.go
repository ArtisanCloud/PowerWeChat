package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerServiceMessageUploadTempMedia struct {
	response.ResponseMiniProgram
	ContentType string `json:"type"`
	MediaID     string `json:"media_id"`
	CreatedAt   int64  `json:"created_at"`
}
