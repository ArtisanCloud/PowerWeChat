package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseWebDriveFileRename struct {
	*response.ResponseWork

	File *power.HashMap `json:"file"`
}
