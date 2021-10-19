package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveFileCreate struct {
	*response.ResponseWork

	FileID string `json:"fileid"`
	Url  string `json:"url"`
}
