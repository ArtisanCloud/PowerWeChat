package request

type RequestAgentSetScope struct {
	AgentID    int      `json:"agentid"`
	AllowUser  []string `json:"allow_user,omitempty"`
	AllowParty []int    `json:"allow_party,omitempty"`
	AllowTag   []int    `json:"allow_tag,omitempty"`
}
