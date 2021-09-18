package request

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
)

type RequestCustomerStrategyEdit struct {
	StrategyID     int64            `json:"strategy_id" `
	StrategyName string           `json:"strategy_name"`
	AdminList    []string         `json:"admin_list"`
	Privilege    *power.HashMap   `json:"privilege"`
	RangeAdd        []*power.HashMap `json:"range_add"`
	RangeDel        []*power.HashMap `json:"range_del"`
}
