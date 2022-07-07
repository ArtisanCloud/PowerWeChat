package request

type RequestTagDelCorpTag struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
	AgentID *int64   `json:"agentid"`
}
