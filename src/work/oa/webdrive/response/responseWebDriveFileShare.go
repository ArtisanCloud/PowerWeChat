package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseWebDriveFileShare struct {
	*response.ResponseWork

	ShareURL string `json:"share_url"`
}
