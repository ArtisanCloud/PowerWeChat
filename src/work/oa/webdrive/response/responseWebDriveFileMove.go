package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveFileMove struct {
	*response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
