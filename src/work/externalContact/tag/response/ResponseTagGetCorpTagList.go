package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseTagGetCorpTagList struct {
	*response.ResponseWork

	ExternalContactList *power.HashMap `json:"tag_group"`
}
