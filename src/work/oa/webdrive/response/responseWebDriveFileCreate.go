package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseWebDriveFileCreate struct {
	*response.ResponseWork

	FileID string `json:"fileid"`
	Url    string `json:"url"`
}
