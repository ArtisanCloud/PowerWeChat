package request

type RequestCustomerStrategyEdit struct {
	StrategyID   int64                            `json:"strategy_id" `
	StrategyName string                           `json:"strategy_name"`
	AdminList    []string                         `json:"admin_list"`
	Privilege    RequestCustomerStrategyPrivilege `json:"privilege"`
	RangeAdd     []RequestCustomerStrategyRange   `json:"range_add"`
	RangeDel     []RequestCustomerStrategyRange   `json:"range_del"`
}
