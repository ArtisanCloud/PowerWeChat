package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveFileShare struct {
	*response.ResponseWork

	ShareURL string `json:"share_url"`
}
