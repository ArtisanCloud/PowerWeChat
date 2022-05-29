package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAgentList struct {
	*response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
