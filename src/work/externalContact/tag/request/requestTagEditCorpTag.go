package request

type RequestTagEditCorpTag struct {
	GroupID   string           `json:"group_id"`
	GroupName string           `json:"group_name"`
	Order     int              `json:"order"`
	AgentID   int64            `json:"agentid"`
}
