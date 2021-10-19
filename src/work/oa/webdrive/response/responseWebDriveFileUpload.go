package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	*response.ResponseWork

	FileID   string             `json:"fileid"`
}
