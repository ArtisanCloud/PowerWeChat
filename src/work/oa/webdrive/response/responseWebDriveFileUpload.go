package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	*response.ResponseWork

	FileID   string             `json:"fileid"`
}
