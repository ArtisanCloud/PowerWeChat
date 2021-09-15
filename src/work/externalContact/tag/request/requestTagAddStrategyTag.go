package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestTagAddStrategyTag struct {
	StrategyID int64            `json:"strategy_id"`
	GroupID    string           `json:"group_id"`
	GroupName  string           `json:"group_name"`
	Order      int              `json:"order"`
	Tag        []*power.HashMap `json:"tag"`
}
