package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveFileMove struct {
	*response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
