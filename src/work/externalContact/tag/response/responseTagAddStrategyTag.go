package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type StrategyTagGroup struct {
	GroupId    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	StrategyID int    `json:"strategy_id"`
	Tag        []*Tag `json:"tag"`
}

type ResponseTagAddStrategyTag struct {
	response.ResponseWork

	TagGroups *StrategyTagGroup `json:"tag_group"`
}
