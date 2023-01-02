package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseAgentList struct {
	response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
