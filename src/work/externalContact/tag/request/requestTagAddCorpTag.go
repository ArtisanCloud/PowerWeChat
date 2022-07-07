package request

type RequestTagAddCorpTag struct {
	GroupID   *string                        `json:"group_id"`
	GroupName string                         `json:"group_name"`
	Order     int                            `json:"order"`
	Tag       []RequestTagAddCorpTagFieldTag `json:"tag"`
	AgentID   *int64                         `json:"agentid"`
}

type RequestTagAddCorpTagFieldTag struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}
