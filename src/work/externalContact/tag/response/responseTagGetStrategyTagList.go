package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseTagGetStrategyTagList struct {
	response.ResponseWork

	TagGroups []*StrategyTagGroup `json:"tag_group"`
}
