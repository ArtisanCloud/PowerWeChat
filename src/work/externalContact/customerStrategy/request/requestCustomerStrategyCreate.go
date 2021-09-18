package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type RequestCustomerStrategyCreate struct {
	ParentID     int64            `json:"parent_id" `
	StrategyName string           `json:"strategy_name"`
	AdminList    []string         `json:"admin_list"`
	Privilege    *power.HashMap   `json:"privilege"`
	Range        []*power.HashMap `json:"range"`
}
