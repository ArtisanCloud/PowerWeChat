package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/agent/request"
)

type ResponseAgentGetWorkbenchTemplate struct {
	*response.ResponseWork

	TemplateType    string                 `json:"type"`
	Image           request.WorkBenchImage `json:"image"`
	ReplaceUserData bool                   `json:"replace_user_data"`
}
