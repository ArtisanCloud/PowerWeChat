package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveFileList struct {
	*response.ResponseWork

	HasMore   bool             `json:"has_more"`
	NextStart int              `json:"next_start"`
	FileList  []*power.HashMap `json:"file_list"`
}
