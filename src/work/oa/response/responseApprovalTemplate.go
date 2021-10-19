package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseApprovalTemplate struct {
	*response.ResponseWork
	TemplateNames   []*power.StringMap `json:"template_names"`
	TemplateContent []*power.HashMap   `json:"template_content"`
}
