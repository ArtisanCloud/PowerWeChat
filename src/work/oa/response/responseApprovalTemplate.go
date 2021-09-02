package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseApprovalTemplate struct {
	*response.ResponseWork
	TemplateNames   []*object.StringMap `json:"template_names"`
	TemplateContent []*object.HashMap   `json:"template_content"`
}
