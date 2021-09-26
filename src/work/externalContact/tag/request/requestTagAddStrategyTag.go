package request

type RequestTagAddStrategyTag struct {
	StrategyID int64                              `json:"strategy_id"`
	GroupID    string                             `json:"group_id"`
	GroupName  string                             `json:"group_name"`
	Order      int                                `json:"order"`
	Tag        []RequestTagAddStrategyTagFieldTag `json:"tag"`
}

type RequestTagAddStrategyTagFieldTag struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}
