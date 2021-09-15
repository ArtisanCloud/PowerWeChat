package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestTagAddCorpTag struct {
	GroupID   string           `json:"group_id"`
	GroupName string           `json:"group_name"`
	Order     int              `json:"order"`
	Tag       []*power.HashMap `json:"tag"`
	AgentID   int64            `json:"agentid"`
}
