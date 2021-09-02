package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAgentGetWorkbenchTemplate struct {
	*response.ResponseWork

	TemplateType    string           `json:"type"`
	Image           power.StringMap `json:"image"`
	ReplaceUserData bool             `json:"replace_user_data"`
}
