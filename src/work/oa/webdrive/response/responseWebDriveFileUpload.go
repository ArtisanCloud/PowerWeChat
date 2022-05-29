package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	*response.ResponseWork

	FileID string `json:"fileid"`
}
