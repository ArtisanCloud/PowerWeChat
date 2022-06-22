package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseTagGetCorpTagList struct {
	response.ResponseWork

	TagGroups []*CorpTagGroup `json:"tag_group"`
}
