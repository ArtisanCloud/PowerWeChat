package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveFileRename struct {
	*response.ResponseWork

	File *power.HashMap `json:"file"`
}
