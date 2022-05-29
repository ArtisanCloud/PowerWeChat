package request

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"

type RequestMomentStrategyCreate struct {
	ParentID     int              `json:"parent_id"`
	StrategyName string           `json:"strategy_name"`
	AdminList    []string         `json:"admin_list"`
	Privilege    *power.HashMap   `json:"privilege"`
	Range        []*power.HashMap `json:"range"`
}
