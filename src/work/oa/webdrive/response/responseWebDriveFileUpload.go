package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	*response.ResponseWork

	FileID   string             `json:"fileid"`
}
