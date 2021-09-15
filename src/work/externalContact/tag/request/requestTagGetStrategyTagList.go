package request

type RequestTagGetStrategyTagList struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
	StrategyID int64    `json:"strategy_id"`
}
