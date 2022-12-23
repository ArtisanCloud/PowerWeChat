package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestMomentStrategyEdit struct {
	StrategyID   int64            `json:"strategy_id"`
	StrategyName string           `json:"strategy_name"`
	AdminList    []string         `json:"admin_list"`
	Privilege    *power.HashMap   `json:"privilege"`
	RangeAdd     []*power.HashMap `json:"range_add"`
	RangeDel     []*power.HashMap `json:"range_del"`
}
