package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseUploadMedia struct {
	response.ResponseWork

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type ResponseUploadMediaData struct {
	response.ResponseWork

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}
