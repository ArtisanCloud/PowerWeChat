package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveFileRename struct {
	*response.ResponseWork

	File *power.HashMap `json:"file"`
}
