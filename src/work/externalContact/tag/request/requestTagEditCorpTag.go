package request

type RequestTagEditCorpTag struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	AgentID *int64 `json:"agentid"`
}
