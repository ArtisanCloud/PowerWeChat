package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAgentList struct {
	*response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
