package request

type RequestTagEditCorpTag struct {
	ID   string           `json:"id"`
	GroupName string           `json:"group_name"`
	Order     int              `json:"order"`
	AgentID   int64            `json:"agentid"`
}
