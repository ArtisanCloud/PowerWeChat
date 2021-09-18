package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveFileCreate struct {
	*response.ResponseWork

	FileID string `json:"fileid"`
	Url  string `json:"url"`
}
