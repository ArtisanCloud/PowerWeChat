package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAgentList struct {
	*response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
