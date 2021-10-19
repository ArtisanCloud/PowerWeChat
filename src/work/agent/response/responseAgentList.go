package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAgentList struct {
	*response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
